package service

import "goweb/model"

type AuthService interface {
	Login(model.LoginRequest) (*model.LoginResponse, error)
	Logout() error
	RegisterUser(param model.RegisterUser) (*string, error)
}
