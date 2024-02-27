package server

import (
	"context"
	"github.com/ahghazey/rate_limiter/pkg/config"
	"net/http"
)

type HttpServer struct {
	server  http.Server
	address string
}

func NewHttpServer(config *config.Config, router http.Handler) *HttpServer {
	return &HttpServer{
		address: config.Address(),
		server: http.Server{
			Addr:           config.Address(),
			Handler:        router,
			ReadTimeout:    config.ReadTimeOut(),
			WriteTimeout:   config.WriteTimeOut(),
			IdleTimeout:    config.IdleTimeOut(),
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *HttpServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
