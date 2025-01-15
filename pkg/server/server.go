package server

import (
	"LoggingApp/internal/config"
	"LoggingApp/pkg/HttpServer"
	"LoggingApp/pkg/handler"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer HttpServer.Server
	httpNet    http.Server
}

func (s *Server) HTTPServer(port string, router *gin.Engine) error {
	_, err := s.httpServer.HTTPServer(port, router)

	//s = &Server{httpServer: s.httpServer, httpNet: *httpServer}

	return err
}
func (s *Server) Run(config *config.Config, h *handler.Handler) error {

	// init server add addr and router
	err := s.HTTPServer(config.BindAddr, h.Routing())

	return err
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpNet.Shutdown(ctx)
}
