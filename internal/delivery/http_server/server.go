package http_server

import (
	"context"
	"fmt"
	"go-shortener/internal/config"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(httpServerConfig config.HttpServerConfig, handler *Handler) *Server {
	server := &http.Server{
		Addr:    httpServerConfig.ListenAddr,
		Handler: handler.engine,
	}

	return &Server{
		server: server,
	}
}

func (s *Server) Run() {
	_ = s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("started server shutdown")
	defer log.Println("finished server shutdown")
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server: %w", s.server.Shutdown(ctx))
	}
	return nil
}
