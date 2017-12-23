package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
	"sync"

	"gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/repo/config"
	"gx/ipfs/QmP9vZfc5WSjfGTXmwX2EcicMFzmZ6fXn7HTdKYat6ccmH/go-ipfs-cmds"
	"gx/ipfs/QmPG2kW5t27LuHgHnvhUwbHCNHAt2eUcb4gPHqofrESUdB/cors"
	logging "gx/ipfs/QmSpJByNKFX1sCsHBEp3R73FL4NF6FnQTEGyNAXHm2GS52/go-log"
	"gx/ipfs/QmSvcDkiRwB8LuMhUtnvhum2C851Mproo75ZDD19jx43tD/go-libp2p-loggables"
)

var log = logging.Logger("cmds/http")

// the internal handler for the API
type internalHandler struct {
	ctx  cmds.Context
	root *cmds.Command
	cfg  *ServerConfig
}

// The Handler struct is funny because we want to wrap our internal handler
// with CORS while keeping our fields.
type Handler struct {
	internalHandler
	corsHandler http.Handler
}

var (
	ErrNotFound           = errors.New("404 page not found")
	errApiVersionMismatch = errors.New("api version mismatch")
)

const (
	StreamErrHeader          = "X-Stream-Error"
	streamHeader             = "X-Stream-Output"
	channelHeader            = "X-Chunked-Output"
	extraContentLengthHeader = "X-Content-Length"
	uaHeader                 = "User-Agent"
	contentTypeHeader        = "Content-Type"
	contentDispHeader        = "Content-Disposition"
	transferEncodingHeader   = "Transfer-Encoding"
	applicationJson          = "application/json"
	applicationOctetStream   = "application/octet-stream"
	plainText                = "text/plain"
	originHeader             = "origin"
)

var AllowedExposedHeadersArr = []string{streamHeader, channelHeader, extraContentLengthHeader}
var AllowedExposedHeaders = strings.Join(AllowedExposedHeadersArr, ", ")

const (
	ACAOrigin      = "Access-Control-Allow-Origin"
	ACAMethods     = "Access-Control-Allow-Methods"
	ACACredentials = "Access-Control-Allow-Credentials"
)

var mimeTypes = map[cmds.EncodingType]string{
	cmds.Protobuf: "application/protobuf",
	cmds.JSON:     "application/json",
	cmds.XML:      "application/xml",
	cmds.Text:     "text/plain",
}

type ServerConfig struct {
	// Headers is an optional map of headers that is written out.
	Headers map[string][]string

	// cORSOpts is a set of options for CORS headers.
	cORSOpts *cors.Options

	// cORSOptsRWMutex is a RWMutex for read/write CORSOpts
	cORSOptsRWMutex sync.RWMutex
}

func skipAPIHeader(h string) bool {
	switch h {
	case "Access-Control-Allow-Origin":
		return true
	case "Access-Control-Allow-Methods":
		return true
	case "Access-Control-Allow-Credentials":
		return true
	default:
		return false
	}
}

func NewHandler(ctx cmds.Context, root *cmds.Command, cfg *ServerConfig) http.Handler {
	if cfg == nil {
		panic("must provide a valid ServerConfig")
	}

	// setup request logger
	ctx.ReqLog = new(cmds.ReqLog)

	// Wrap the internal handler with CORS handling-middleware.
	// Create a handler for the API.
	internal := internalHandler{
		ctx:  ctx,
		root: root,
		cfg:  cfg,
	}
	c := cors.New(*cfg.cORSOpts)
	return &Handler{internal, c.Handler(internal)}
}

func (i Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call the CORS handler which wraps the internal handler.
	i.corsHandler.ServeHTTP(w, r)
}

func (i internalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("incoming API request: ", r.URL)

	defer func() {
		if r := recover(); r != nil {
			log.Error("a panic has occurred in the commands handler!")
			log.Error(r)
			log.Errorf("stack trace:\n%s", debug.Stack())
		}
	}()

	// get the node's context to pass into the commands.
	node, err := i.ctx.GetNode()
	if err != nil {
		s := fmt.Sprintf("cmds/http: couldn't GetNode(): %s", err)
		http.Error(w, s, http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithCancel(node.Context())
	defer cancel()
	ctx = logging.ContextWithLoggable(ctx, loggables.Uuid("requestId"))
	if cn, ok := w.(http.CloseNotifier); ok {
		clientGone := cn.CloseNotify()
		go func() {
			select {
			case <-clientGone:
			case <-ctx.Done():
			}
			cancel()
		}()
	}

	if !allowOrigin(r, i.cfg) || !allowReferer(r, i.cfg) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 - Forbidden"))
		log.Warningf("API blocked request to %s. (possible CSRF)", r.URL)
		return
	}

	req, err := Parse(r, i.root)
	if err != nil {
		if err == ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(err.Error()))
		return
	}

	reqLogEnt := i.ctx.ReqLog.Add(req)
	defer i.ctx.ReqLog.Finish(reqLogEnt)

	//ps: take note of the name clash - commands.Context != context.Context
	req.SetInvocContext(i.ctx)

	err = req.SetRootContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set user's headers first.
	for k, v := range i.cfg.Headers {
		if !skipAPIHeader(k) {
			w.Header()[k] = v
		}
	}

	re := NewResponseEmitter(w, r.Method, req)
	defer re.Close()

	// call the command
	err = i.root.Call(req, re)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := re.(Doner)
	select {
	case <-d.Done():
	case <-req.Context().Done():
		log.Errorf("waiting for http.Responseemitter to close but then %s", req.Context().Err())
		// too late to send an error, just return
	}

	return
}

func sanitizedErrStr(err error) string {
	s := err.Error()
	s = strings.Split(s, "\n")[0]
	s = strings.Split(s, "\r")[0]
	return s
}

func NewServerConfig() *ServerConfig {
	cfg := new(ServerConfig)
	cfg.cORSOpts = new(cors.Options)
	return cfg
}

func (cfg *ServerConfig) AllowedOrigins() []string {
	cfg.cORSOptsRWMutex.RLock()
	defer cfg.cORSOptsRWMutex.RUnlock()
	return cfg.cORSOpts.AllowedOrigins
}

func (cfg *ServerConfig) SetAllowedOrigins(origins ...string) {
	cfg.cORSOptsRWMutex.Lock()
	defer cfg.cORSOptsRWMutex.Unlock()
	o := make([]string, len(origins))
	copy(o, origins)
	cfg.cORSOpts.AllowedOrigins = o
}

func (cfg *ServerConfig) AppendAllowedOrigins(origins ...string) {
	cfg.cORSOptsRWMutex.Lock()
	defer cfg.cORSOptsRWMutex.Unlock()
	cfg.cORSOpts.AllowedOrigins = append(cfg.cORSOpts.AllowedOrigins, origins...)
}

func (cfg *ServerConfig) AllowedMethods() []string {
	cfg.cORSOptsRWMutex.RLock()
	defer cfg.cORSOptsRWMutex.RUnlock()
	return []string(cfg.cORSOpts.AllowedMethods)
}

func (cfg *ServerConfig) SetAllowedMethods(methods ...string) {
	cfg.cORSOptsRWMutex.Lock()
	defer cfg.cORSOptsRWMutex.Unlock()
	if cfg.cORSOpts == nil {
		cfg.cORSOpts = new(cors.Options)
	}
	cfg.cORSOpts.AllowedMethods = methods
}

func (cfg *ServerConfig) SetAllowCredentials(flag bool) {
	cfg.cORSOptsRWMutex.Lock()
	defer cfg.cORSOptsRWMutex.Unlock()
	cfg.cORSOpts.AllowCredentials = flag
}

// allowOrigin just stops the request if the origin is not allowed.
// the CORS middleware apparently does not do this for us...
func allowOrigin(r *http.Request, cfg *ServerConfig) bool {
	origin := r.Header.Get("Origin")

	// curl, or ipfs shell, typing it in manually, or clicking link
	// NOT in a browser. this opens up a hole. we should close it,
	// but right now it would break things. TODO
	if origin == "" {
		return true
	}
	origins := cfg.AllowedOrigins()
	for _, o := range origins {
		if o == "*" { // ok! you asked for it!
			return true
		}

		if o == origin { // allowed explicitly
			return true
		}
	}

	return false
}

// allowReferer this is here to prevent some CSRF attacks that
// the API would be vulnerable to. We check that the Referer
// is allowed by CORS Origin (origins and referrers here will
// work similarly in the normla uses of the API).
// See discussion at https://github.com/ipfs/go-ipfs/issues/1532
func allowReferer(r *http.Request, cfg *ServerConfig) bool {
	referer := r.Referer()

	// curl, or ipfs shell, typing it in manually, or clicking link
	// NOT in a browser. this opens up a hole. we should close it,
	// but right now it would break things. TODO
	if referer == "" {
		return true
	}

	u, err := url.Parse(referer)
	if err != nil {
		// bad referer. but there _is_ something, so bail.
		// debug because referer comes straight from the client. dont want to
		// let people DOS by putting a huge referer that gets stored in log files.
		return false
	}
	origin := u.Scheme + "://" + u.Host

	// check CORS ACAOs and pretend Referer works like an origin.
	// this is valid for many (most?) sane uses of the API in
	// other applications, and will have the desired effect.
	origins := cfg.AllowedOrigins()
	for _, o := range origins {
		if o == "*" { // ok! you asked for it!
			return true
		}

		// referer is allowed explicitly
		if o == origin {
			return true
		}
	}

	return false
}

// apiVersionMatches checks whether the api client is running the
// same version of go-ipfs. for now, only the exact same version of
// client + server work. In the future, we should use semver for
// proper API versioning! \o/
func apiVersionMatches(r *http.Request) error {
	clientVersion := r.UserAgent()
	// skips check if client is not go-ipfs
	if clientVersion == "" || !strings.Contains(clientVersion, "/go-ipfs/") {
		return nil
	}

	daemonVersion := config.ApiVersion
	if daemonVersion != clientVersion {
		return fmt.Errorf("%s (%s != %s)", errApiVersionMismatch, daemonVersion, clientVersion)
	}
	return nil
}