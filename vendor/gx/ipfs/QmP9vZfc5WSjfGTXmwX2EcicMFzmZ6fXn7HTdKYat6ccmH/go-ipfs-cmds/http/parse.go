package http

import (
	"errors"
	"fmt"
	"mime"
	"net/http"
	"strings"

	path "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/path"
	cmds "gx/ipfs/QmP9vZfc5WSjfGTXmwX2EcicMFzmZ6fXn7HTdKYat6ccmH/go-ipfs-cmds"
	"gx/ipfs/QmQp2a2Hhb7F6eK2A5hN8f9aJy4mtkEikL9Zj4cgB7d1dD/go-ipfs-cmdkit"
	"gx/ipfs/QmQp2a2Hhb7F6eK2A5hN8f9aJy4mtkEikL9Zj4cgB7d1dD/go-ipfs-cmdkit/files"
)

// Parse parses the data in a http.Request and returns a command Request object
func Parse(r *http.Request, root *cmds.Command) (cmds.Request, error) {
	if !strings.HasPrefix(r.URL.Path, ApiPath) {
		return nil, errors.New("Unexpected path prefix")
	}
	pth := path.SplitList(strings.TrimPrefix(r.URL.Path, ApiPath+"/"))

	stringArgs := make([]string, 0)

	if err := apiVersionMatches(r); err != nil {
		if pth[0] != "version" { // compatibility with previous version check
			return nil, err
		}
	}

	getPath := pth[:len(pth)-1]
	cmd, err := root.Get(getPath)
	if err != nil {
		// 404 if there is no command at that path
		return nil, ErrNotFound

	}

	sub := cmd.Subcommand(pth[len(pth)-1])

	if sub == nil {
		if cmd.Run == nil {
			return nil, ErrNotFound
		}

		// if the last string in the path isn't a subcommand, use it as an argument
		// e.g. /objects/Qabc12345 (we are passing "Qabc12345" to the "objects" command)
		stringArgs = append(stringArgs, pth[len(pth)-1])
		pth = pth[:len(pth)-1]

	} else {
		cmd = sub
	}

	opts, stringArgs2 := parseOptions(r)
	stringArgs = append(stringArgs, stringArgs2...)

	// count required argument definitions
	numRequired := 0
	for _, argDef := range cmd.Arguments {
		if argDef.Required {
			numRequired++
		}
	}

	// count the number of provided argument values
	valCount := len(stringArgs)

	args := make([]string, valCount)

	valIndex := 0
	requiredFile := ""
	for _, argDef := range cmd.Arguments {
		// skip optional argument definitions if there aren't sufficient remaining values
		if valCount-valIndex <= numRequired && !argDef.Required {
			continue
		} else if argDef.Required {
			numRequired--
		}

		if argDef.Type == cmdkit.ArgString {
			if argDef.Variadic {
				for _, s := range stringArgs {
					args[valIndex] = s
					valIndex++
				}
				valCount -= len(stringArgs)

			} else if len(stringArgs) > 0 {
				args[valIndex] = stringArgs[0]
				stringArgs = stringArgs[1:]
				valIndex++

			} else {
				break
			}
		} else if argDef.Type == cmdkit.ArgFile && argDef.Required && len(requiredFile) == 0 {
			requiredFile = argDef.Name
		}
	}

	optDefs, err := root.GetOptions(pth)
	if err != nil {
		return nil, err
	}

	// create cmds.File from multipart/form-data contents
	contentType := r.Header.Get(contentTypeHeader)
	mediatype, _, _ := mime.ParseMediaType(contentType)

	var f files.File
	if mediatype == "multipart/form-data" {
		reader, err := r.MultipartReader()
		if err != nil {
			return nil, err
		}

		f = &files.MultipartFile{
			Mediatype: mediatype,
			Reader:    reader,
		}
	}

	// if there is a required filearg, error if no files were provided
	if len(requiredFile) > 0 && f == nil {
		return nil, fmt.Errorf("File argument '%s' is required", requiredFile)
	}

	req, err := cmds.NewRequest(pth, opts, args, f, cmd, optDefs)
	if err != nil {
		return nil, err
	}

	err = cmd.CheckArguments(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func parseOptions(r *http.Request) (map[string]interface{}, []string) {
	opts := make(map[string]interface{})
	var args []string

	query := r.URL.Query()
	for k, v := range query {
		if k == "arg" {
			args = v
		} else {
			opts[k] = v[0]
		}
	}

	// default to setting encoding to JSON
	_, short := opts[cmdkit.EncShort]
	_, long := opts[cmdkit.EncLong]
	if !short && !long {
		opts[cmdkit.EncShort] = cmds.JSON
	}

	return opts, args
}