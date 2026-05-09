package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) mount() *chi.Mux {
	// mux := http.NewServeMux()

	// mux.HandleFunc("GET /v1/health", app.healthcheck)

	r := chi.NewRouter()

	//middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)    //adds loging information such as request status etc
	r.Use(middleware.Recoverer) //recovers from panic condition

	r.Get("/v1/health", app.healthcheck)
	r.Get("/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, //if the server takes more than 30 sec's to write a response to the client then its going to timeOut
		ReadTimeout:  time.Second * 10, //if the client takes more than 10 sec's to read the response from the server then its goind to timeOut
		IdleTimeout:  time.Minute,
	}

	log.Printf("server started at %s", app.config.addr)

	return srv.ListenAndServe()
}
