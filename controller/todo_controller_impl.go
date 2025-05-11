package controller

import (
	"goweb/model"
	"goweb/service"
	"goweb/utils"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	todoService service.TodoService
}

// ModifyTodo implements TodoController.
func (ctrl *TodoControllerImpl) ModifyTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIDRaw := request.Context().Value("user_id")

	if userIDRaw == nil {
		utils.ResponseBody(writer, 401, "Unauthorized", "Who are you?")
		return
	}

	userID := userIDRaw.(string)
	var updateReq model.Todo
	todoId := params.ByName("id")

	if todoId == "" {
		utils.ResponseBody(writer, http.StatusBadRequest, "Bad Request", "Missing ID")
		return
	}

	if !utils.RequestBody(writer, request, &updateReq) {
		return
	}

	updateReq.ModifyAt = time.Now()
	updateReq.UserId = userID

	response, err := ctrl.todoService.Modify(todoId, updateReq)
	if err != nil {
		utils.ResponseBody(writer, 200, "OK", err.Error())
		return
	}

	utils.ResponseBody(writer, 200, "OK", &response)
	return

}

// CreateTodo implements TodoController.
// Todo parameter need Title and Content, and the rest value created by backend
func (ctrl *TodoControllerImpl) CreateTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	/// Get user id from context, since create todo needed for user login
	userIDRaw := request.Context().Value("user_id")

	if userIDRaw == nil || userIDRaw == "" {
		utils.ResponseBody(writer, 401, "Unauthorized", "Who are you?")
		return
	}

	userID := userIDRaw.(string)
	var createReq model.Todo
	if !utils.RequestBody(writer, request, &createReq) {
		return
	}
	createReq.UserId = userID
	createReq.CreatedAt = time.Now()
	createReq.ModifyAt = time.Now()

	response, err := ctrl.todoService.Create(createReq)
	if err != nil {
		utils.ResponseBody(writer, http.StatusOK, "OK", err.Error())
		return
	}
	utils.ResponseBody(writer, http.StatusCreated, "Created", &response)
	return
}

// DeleteTodo implements TodoController.
func (ctrl *TodoControllerImpl) DeleteTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	if todoId == "" {
		utils.ResponseBody(writer, http.StatusBadRequest, "Bad Requests", "Missing ID")
		return
	}

	response, err := ctrl.todoService.Delete(todoId)

	if err != nil {
		utils.ResponseBody(writer, http.StatusOK, "OK", err.Error())
		return
	}

	utils.ResponseBody(writer, http.StatusOK, "OK", &response)
	return
}

// ListTodo implements TodoController.
func (ctrl *TodoControllerImpl) ListTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIDRaw := request.Context().Value("user_id")

	if userIDRaw == nil || userIDRaw == "" {
		utils.ResponseBody(writer, 401, "Unauthorized", "Who are you?")
		return
	}

	userID := userIDRaw.(string)

	response, err := ctrl.todoService.List(userID)

	if err != nil {
		utils.ResponseBody(writer, http.StatusOK, "OK", err.Error())
		return
	}

	utils.ResponseBody(writer, http.StatusOK, "OK", &response)
	return
}

// SingleTodo implements TodoController.
func (ctrl *TodoControllerImpl) SingleTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todoId := params.ByName("id")

	if todoId == "" {
		utils.ResponseBody(writer, http.StatusBadRequest, "Bad Requests", "Missing ID")
		return
	}

	response, err := ctrl.todoService.SingleTodo(todoId)

	if err != nil {
		utils.ResponseBody(writer, http.StatusOK, "OK", err.Error())
		return
	}

	utils.ResponseBody(writer, http.StatusOK, "OK", &response)
	return

}

func NewTodoController(service service.TodoService) TodoController {
	return &TodoControllerImpl{
		todoService: service,
	}
}
