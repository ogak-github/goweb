package repository

import (
	"context"
	"errors"
	"fmt"
	"goweb/model"
	"goweb/utils"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepositoryImpl struct {
}

// Login implements AuthRepository.
func (service *AuthRepositoryImpl) Login(db *pgxpool.Pool, param model.LoginRequest) (*model.LoginResponse, error) {
	loginQuery := `SELECT id, password FROM users WHERE username = $1`

	var encryptedPassword string
	var userId string
	err := db.QueryRow(context.Background(), loginQuery, param.Username).Scan(&userId, &encryptedPassword)

	if err != nil {
		fmt.Println("Database error: " + err.Error())
		return nil, err
	}

	if !utils.VerifyHash(param.Password, encryptedPassword) {
		return nil, errors.New("Invalid username or password")
	}

	var generatedJWT string

	generatedJWT, err = utils.GenerateJWT(userId)
	if err != nil {
		fmt.Println("Failed to generate JWT: " + err.Error()) // log for dev
		return nil, fmt.Errorf(err.Error())
	}

	var result = model.LoginResponse{
		Token:     generatedJWT,
		ExpiredIn: time.Now().Add(24 * time.Hour).Local().Format(time.UnixDate),
	}

	return &result, nil

}

// Logout implements AuthRepository.
func (a *AuthRepositoryImpl) Logout() error {
	panic("unimplemented")
}

// RegisterUser implements AuthRepository.
func (a *AuthRepositoryImpl) RegisterUser(db *pgxpool.Pool, register model.RegisterUser) (*string, error) {

	registerQuery := `INSERT INTO users (username, password, full_name, email) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(context.Background(), registerQuery, 
		register.Username,
		register.Password,
		register.FullName,
		register.Email)

	if err != nil {
		fmt.Println("Database error: " + err.Error())
		return nil, fmt.Errorf("Database error, please contact administrator")
	}
	var result = register.Username + " has been registered, you'll need to verify your email"

	return &result, nil
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}
