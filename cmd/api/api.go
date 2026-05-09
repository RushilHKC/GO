package main

import (
	"log"
	"net/http"
	"time"
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

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthcheck)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

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
