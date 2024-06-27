package server

import (
	"log"
	"net/http"
)

type Server struct {
	http *http.ServeMux
}

func Init() Server {
	mux := http.NewServeMux()
	return Server{http: mux}
}

func (s Server) Start() {
	s.http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	})
	log.Fatal(http.ListenAndServe(":8080", s.http))
}
