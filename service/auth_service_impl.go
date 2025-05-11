package service

import (
	"errors"
	"goweb/model"
	"goweb/repository"
	"goweb/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authServiceImpl struct {
	AuthRepository repository.AuthRepository
	Db             *pgxpool.Pool
	v              *validator.Validate
	///DB, Repos, or anything needed
}

// Login implements AuthService.
func (service *authServiceImpl) Login(params model.LoginRequest) (*model.LoginResponse, error) {
	err := service.v.Struct(params)
	if err != nil {
		return nil, errors.New(utils.FormatValidationError(err))
	}
	return service.AuthRepository.Login(service.Db, params)
}

// Logout implements AuthService.
func (a *authServiceImpl) Logout() error {
	panic("unimplemented")
}

// RegisterUser implements AuthService.
func (service *authServiceImpl) RegisterUser(param model.RegisterUser) (*string, error) {
	err := service.v.Struct(param)

	if err != nil {
		return nil, errors.New(utils.FormatValidationError(err))
	}
	param.Password = utils.HashPassword(param.Password)
	return service.AuthRepository.RegisterUser(service.Db, param)
}

func NewAuthService(db *pgxpool.Pool, repo repository.AuthRepository, newValidator *validator.Validate) AuthService {
	return &authServiceImpl{
		AuthRepository: repo,
		Db:             db,
		v:              newValidator,
	}
}
