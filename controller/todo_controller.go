package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	CreateTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ListTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AllTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SingleTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ModifyTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
