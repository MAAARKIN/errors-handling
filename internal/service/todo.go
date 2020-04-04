package service

import (
	"github.com/maaarkin/errors-handling/pkg/domains/todo"
	"github.com/pkg/errors"
)

//TodoService representa o servico que poderá ser exporto para a camada de consumo (rest, grpc, rabbitmq, etc...)
type TodoService interface {
	Save(title string) (*todo.Todo, error)
}

type todoService struct {
}

//NewTodoService função que tem como responsabilidade inicializar a interface do TodoService
func NewTodoService() TodoService {
	return todoService{}
}

func (t todoService) Save(value string) (*todo.Todo, error) {
	if value == "" {
		return nil, errors.WithStack(todo.TodoTitleError{})
	}

	if value == "teste" {
		return nil, todo.TodoTitleError{}
	}

	out := todo.Todo{
		Title: value,
	}

	return &out, nil
}
