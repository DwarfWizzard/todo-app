package main

import (
	"log"

	todoapp "github.com/DwarfWizzard/todo-app"
	"github.com/DwarfWizzard/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	
	srv := new(todoapp.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}