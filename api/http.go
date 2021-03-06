package api

import (
	"fmt"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/iain17/decentralizer/pb"
	"github.com/iain17/logger"
	"net/http"
	_ "net/http/pprof"
)

func (s *Server) initHTTP(port int) error {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	mux := runtime.NewServeMux()
	address := fmt.Sprintf(":%d", port)
	logger.Infof("Serving HTTP API on: %s", address)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterDecentralizerHandlerFromEndpoint(s.ctx, mux, s.endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(address, mux)
}
