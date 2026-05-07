package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
	api := &api{":8080"}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	server.ListenAndServe()
}
