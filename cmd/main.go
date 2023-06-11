package main

import (
	"log"

	"github.com/MarselBisengaliev/go-todo-app"
	handler "github.com/MarselBisengaliev/go-todo-app/cmd/pkg/handlers"
)

func main() {
	handlers := new(handler.Handler)
	
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occur	ed while running http server: %s", err.Error())
	}
}
