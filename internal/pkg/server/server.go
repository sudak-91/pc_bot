package server

import (
	"bytes"
	"log"
	"net/http"

	update "github.com/sudak-91/telegrambotgo/Service"
)

type Server struct {
	port    string
	key     string
	updater update.Updater
}

func NewServer(Port string, key string, Upd update.Updater) *Server {
	return &Server{
		port:    Port,
		key:     key,
		updater: Upd,
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/"+s.key, s.Handl)
	log.Fatal(http.ListenAndServe(":"+s.port, mux))
}

func (s *Server) Handl(w http.ResponseWriter, r *http.Request) {
	var b []byte
	buffer := bytes.NewBuffer(b)
	buffer.ReadFrom(r.Body)
	k, err := s.updater.Update(buffer.Bytes())
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json")
	_, err = w.Write(k)
	if err != nil {
		panic(err.Error())
	}
}
