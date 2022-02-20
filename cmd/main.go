package main

import (
	"log"

	todoapp "github.com/DwarfWizzard/todo-app"
	"github.com/DwarfWizzard/todo-app/pkg/handler"
	"github.com/DwarfWizzard/todo-app/pkg/repository"
	"github.com/DwarfWizzard/todo-app/pkg/service"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host: "localhost",
		Port: "5432",
		Username: "postgres",
		Password: "1234",
		DBName: "todo-db",
		SSLMode: "disable",
	})

	if err != nil {
		log.Panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(todoapp.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}