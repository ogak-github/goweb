package repository

import (
	"goweb/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository interface {
	Create(db *pgxpool.Pool, params model.Todo) (*string, error)
	List(db *pgxpool.Pool, userId string) (*[]model.Todo, error)
	Modify(db *pgxpool.Pool, todoId string, params model.Todo) (*string, error)
	Delete(db *pgxpool.Pool, todoId string) (*string, error)
	SingleData(db *pgxpool.Pool, todoId string) (*model.Todo, error)
}
