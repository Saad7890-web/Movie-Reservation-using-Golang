package server

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}
}

func (s *Server) Start() {
	log.Printf("Server running on %s\n", s.httpServer.Addr)
	log.Fatal(s.httpServer.ListenAndServe())
}
