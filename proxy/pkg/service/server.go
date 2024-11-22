package service

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	IP         string
	Port       string
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.IP = "0.0.0.0"
	s.Port = port
	s.httpServer = &http.Server{
		Addr:         "0.0.0.0:" + port,
		Handler:      handler,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
