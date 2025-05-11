package repository

import (
	"goweb/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	Login(db *pgxpool.Pool, param model.LoginRequest) (*model.LoginResponse, error)
	Logout() error
	RegisterUser(db *pgxpool.Pool, register model.RegisterUser) (*string, error)
}
