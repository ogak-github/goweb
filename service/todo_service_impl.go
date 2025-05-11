package service

import (
	"errors"
	"fmt"
	"goweb/model"
	"goweb/repository"
	"goweb/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoServiceImpl struct {
	db         *pgxpool.Pool
	repository repository.TodoRepository
	v          *validator.Validate
}

// Create implements TodoService.
func (service *TodoServiceImpl) Create(params model.Todo) (*string, error) {
	err := service.v.Struct(params)

	if err != nil {
		return nil, errors.New(utils.FormatValidationError(err))
	}

	return service.repository.Create(service.db, params)
}

// Delete implements TodoService.
func (service *TodoServiceImpl) Delete(todoId string) (*string, error) {
	return service.repository.Delete(service.db, todoId)
}

// List implements TodoService.
func (service *TodoServiceImpl) List(userId string) (*[]model.Todo, error) {
	return service.repository.List(service.db, userId)
}

// Modify implements TodoService.
func (service *TodoServiceImpl) Modify(todoId string, params model.Todo) (*string, error) {
	var err = service.v.Struct(params)
	if err != nil {
		fmt.Println("Error Input: " + err.Error())
		return nil, errors.New(utils.FormatValidationError(err))
	}

	return service.repository.Modify(service.db, todoId, params)

}

// SingleTodo implements TodoService.
func (service *TodoServiceImpl) SingleTodo(todoId string) (*model.Todo, error) {
	return service.repository.SingleData(service.db, todoId)
}

func NewTodoService(db *pgxpool.Pool, repo repository.TodoRepository, newValidator *validator.Validate) TodoService {
	return &TodoServiceImpl{
		db:         db,
		repository: repo,
		v:          newValidator,
	}
}
