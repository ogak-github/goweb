package repository

import (
	"context"
	"errors"
	"fmt"
	"goweb/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepositoryImpl struct {
}

// AllList implements TodoRepository.
func (t TodoRepositoryImpl) AllList(db *pgxpool.Pool) (*[]model.Todo, error) {
	var todos []model.Todo
	getTodoListQuery := `SELECT * FROM todo`

	err := pgxscan.Select(context.Background(), db, &todos, getTodoListQuery)

	if err != nil {
		fmt.Println("Database " + err.Error())
		return nil, err
	}

	return &todos, nil
}

// SingleData implements TodoRepository.
func (t TodoRepositoryImpl) SingleData(db *pgxpool.Pool, todoId string) (*model.Todo, error) {
	var todo model.Todo
	singleSelectQuery := `SELECT * FROM todo WHERE id = $1`
	err := pgxscan.Get(context.Background(), db, &todo, singleSelectQuery, todoId)

	if err != nil {
		fmt.Println("Database " + err.Error())
		return nil, err
	}
	return &todo, nil
}

// Create implements TodoRepository.
func (t TodoRepositoryImpl) Create(db *pgxpool.Pool, params model.Todo) (*string, error) {
	createQuery := `INSERT INTO todo (title, content, created_at, modify_at, user_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(context.Background(), createQuery, params.Title, params.Content, params.CreatedAt, params.ModifyAt, params.UserId)
	if err != nil {
		return nil, errors.New("Database error: " + err.Error())
	}
	var result = "Todo created"
	return &result, nil
}

// Delete implements TodoRepository.
func (t TodoRepositoryImpl) Delete(db *pgxpool.Pool, todoId string) (*string, error) {
	singleDeleteQuery := `DELETE FROM todo WHERE id=$1`
	tag, err := db.Exec(context.Background(), singleDeleteQuery, todoId)
	if err != nil {
		fmt.Println("Database " + err.Error())
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		return nil, fmt.Errorf("No record found to delete")
	}

	var result = "Todo deleted"
	return &result, nil
}

// List implements TodoRepository.
func (t TodoRepositoryImpl) List(db *pgxpool.Pool, userId string) (*[]model.Todo, error) {
	var todos []model.Todo
	getTodoListQuery := `SELECT * FROM todo WHERE user_id=$1`

	err := pgxscan.Select(context.Background(), db, &todos, getTodoListQuery, userId)

	if err != nil {
		fmt.Println("Database " + err.Error())
		return nil, err
	}

	return &todos, nil
}

// Modify implements TodoRepository.
func (t TodoRepositoryImpl) Modify(db *pgxpool.Pool, todoId string, params model.Todo) (*string, error) {
	updateQuery := "UPDATE todo SET title=$2, content=$3, modify_at=$4 WHERE id = $1"

	_, err := db.Exec(context.Background(), updateQuery, todoId, params.Title, params.Content, params.ModifyAt)

	if err != nil {
		fmt.Println("Database " + err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	var result = "Update success"

	return &result, nil
}

func NewTodoRepository() TodoRepository {
	return TodoRepositoryImpl{}
}
