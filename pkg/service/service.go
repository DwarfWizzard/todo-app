package service

import (
	"github.com/DwarfWizzard/todo-app/pkg/models"
	"github.com/DwarfWizzard/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
