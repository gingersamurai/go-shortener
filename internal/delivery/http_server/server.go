package http_server

import "github.com/gin-gonic/gin"

type Server struct {
	host   string
	engine *gin.Engine
}

func NewServer(host string, handler *Handler) *Server {
	engine := gin.Default()

	engine.POST("/shorten", handler.AddLinkHandler)
	engine.GET("/:mapping", handler.GetLinkHandler)

	return &Server{
		host:   host,
		engine: engine,
	}
}

func (s *Server) Run() error {
	return s.engine.Run(s.host)
}
