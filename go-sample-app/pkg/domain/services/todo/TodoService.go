package services

import (
	TodoRepositories "go-sample-app/pkg/domain/repositories/todo"
)

type TodoService struct {
	Rep TodoRepositories.ITodoRepository
}

func NewTodoService(rep TodoRepositories.ITodoRepository) *TodoService {
	userService := new(TodoService)
	userService.Rep = rep

	return userService
}

func (s *TodoService) CreateTodo() {
	s.Rep.NewTodo()
}
