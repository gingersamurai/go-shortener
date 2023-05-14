package http_server

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	host   string
	server *http.Server
}

func NewServer(host string, handler *Handler) *Server {
	server := &http.Server{
		Addr:    host,
		Handler: handler.engine,
	}

	return &Server{
		host:   host,
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
