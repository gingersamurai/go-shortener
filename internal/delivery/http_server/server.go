package http_server

import (
	"context"
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

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("graceful shutdown server")
	return s.server.Shutdown(ctx)
}
