package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	Errs     chan error
}

func New(host string, port string, errs chan error) Server {
	gin.SetMode(gin.ReleaseMode)
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%v", host, port),
		Errs:     errs,
	}

	return srv
}

func (s *Server) Run() {
	go func() {
		log.Println("Server running on", s.httpAddr)
		s.Errs <- s.engine.Run(s.httpAddr)
	}()
}

func (s *Server) RegisterRoute(httpMethod string, relativePath string, handler gin.HandlerFunc) {
	s.engine.RouterGroup.Handle(httpMethod, relativePath, handler)
}
