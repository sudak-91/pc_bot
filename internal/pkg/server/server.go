package server

import (
	"bytes"
	"log"
	"net/http"
	"sync"

	update "github.com/sudak-91/telegrambotgo/Service"
)

var Util *Utl

type Server struct {
	port    string
	key     string
	updater update.Updater
	once    sync.Once
}

func NewServer(Port string, key string, Upd update.Updater) *Server {
	return &Server{
		port:    Port,
		key:     key,
		updater: Upd,
	}
}

type Utl struct {
	Stage map[int64]int
}

func (s *Server) Run() {
	s.once.Do(func() {
		Util = &Utl{Stage: make(map[int64]int)}
	})
	mux := http.NewServeMux()
	mux.HandleFunc("/"+s.key, s.Handl)
	log.Println("Server start")
	log.Fatal(http.ListenAndServe(":"+s.port, mux))
}

func (s *Server) Handl(w http.ResponseWriter, r *http.Request) {
	var b []byte
	log.Println("New request")
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
