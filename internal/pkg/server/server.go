package server

import (
	"log"
	"net/http"
)

type Server struct {
	port    string
	key     string
	handler http.Handler
}

func NewServer(Port string, key string) *Server {
	return &Server{
		port: Port,
		key:  key,
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()
	mux.Handle("/"+s.key, s.handler)
	log.Fatal(http.ListenAndServe(":"+s.port, s.handler))
}
