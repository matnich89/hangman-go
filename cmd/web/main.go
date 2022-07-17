package main

import (
	"github.com/gorilla/mux"
	"hangman/game"
	"hangman/handler"
	"log"
	"sync"
)

func main() {

	store := game.NewGameStore()
	service := game.NewGameService(store)
	handler := handler.NewHandler(service, handler.NewClientConnections())

	app := &app{
		mux:     mux.NewRouter(),
		handler: handler,
		wg:      sync.WaitGroup{},
	}

	go app.handler.BroadCast()
	err := app.serve()

	if err != nil {
		log.Fatalf("could not start server %s", err.Error())
	}
}
