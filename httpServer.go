package main

import (
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
		default:
			w.Write([]byte("invalid request"))
		}
	default:
		w.Write([]byte("page not found"))
	}
}

func main() {
	s := &server{":8080"}
	http.ListenAndServe(s.addr, s)
}
