package grpc_server

import (
	"context"
	"go-shortener/api/link"
	"go-shortener/internal/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	engine *grpc.Server
	lis    net.Listener
}

func NewServer(serverConfig config.GrpcServerConfig, handler *Handler) (*Server, error) {
	s := &Server{}
	lis, err := net.Listen("tcp", serverConfig.ListenAddr)
	if err != nil {
		return nil, err
	}
	s.lis = lis

	s.engine = grpc.NewServer()
	link.RegisterLinkShortenerServiceServer(s.engine, handler)

	return s, nil
}

func (s *Server) Run() error {
	return s.engine.Serve(s.lis)
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("started grpc_server shutdown")
	defer log.Println("finished grpc_server shutdown")
	s.engine.GracefulStop()
	return nil
}
