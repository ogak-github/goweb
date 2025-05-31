package service

import "goweb/model"

type TodoService interface {
	Create(params model.Todo) (*string, error)
	List(userId string) (*[]model.Todo, error)
	AllList() (*[]model.Todo, error)
	SingleTodo(todoId string) (*model.Todo, error)
	Delete(todoId string) (*string, error)
	Modify(todoId string, params model.Todo) (*string, error)
}
