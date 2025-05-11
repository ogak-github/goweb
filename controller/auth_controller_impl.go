package controller

import (
	"goweb/model"
	"goweb/service"
	"goweb/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

// @Summary Login
// @Description login authentication need username and password
// @Tags auth
// @produce json
// @Success 200 {object} model.LoginResponse
// @Router /api/login [post]
func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var loginReq model.LoginRequest
	if !utils.RequestBody(writer, request, &loginReq) {
		return
	}
	response, err := controller.AuthService.Login(loginReq)
	if err != nil {
		utils.ResponseBody(writer, http.StatusOK, "OK", err.Error())
		return
	}
	utils.ResponseBody(writer, http.StatusOK, "OK", &response)
}

// Logout implements AuthController.
func (a *AuthControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// @Summary User Registration
// @Description Create user
// @Tags auth
// @produce json
// @Success 200 {object} string
// @Router /api/register [post]
func (controller *AuthControllerImpl) RegisterUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var registerReq model.RegisterUser
	if !utils.RequestBody(writer, request, &registerReq) {
		return
	}
	response, err := controller.AuthService.RegisterUser(registerReq)
	if err != nil {
		utils.ResponseBody(writer, http.StatusOK, "OK", err)
		return
	}
	utils.ResponseBody(writer, http.StatusCreated, "Created", &response)
	return
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}
