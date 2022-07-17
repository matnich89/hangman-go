package main

import (
	"github.com/gorilla/mux"
	"hangman/handler"
	"net/http"
	"sync"
)

type app struct {
	mux     *mux.Router
	handler *handler.Handler
	wg      sync.WaitGroup
}

func (a *app) routes() {
	a.mux.HandleFunc("/create", a.handler.CreateGame).Methods(http.MethodGet)
	a.mux.HandleFunc("/connect/{id}", a.handler.ConnectGame).Methods(http.MethodGet)
	a.mux.HandleFunc("/guess/{id}", a.handler.MakeGuessForGame).Methods(http.MethodPut)
}
