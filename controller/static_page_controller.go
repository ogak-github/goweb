package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StaticPageController interface {
	IndexPage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	NotFound(writer http.ResponseWriter, request *http.Request)
	Redirect(writer http.ResponseWriter, request *http.Request)
}

type StaticPageControllerImpl struct{}

// IndexPage implements StaticPageController.
func (s *StaticPageControllerImpl) IndexPage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	http.ServeFile(writer, request, "views/index.html")
}

// NotFound implements StaticPageController.
func (s *StaticPageControllerImpl) NotFound(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "views/404.html")
}

// Redirect implements StaticPageController.
func (s *StaticPageControllerImpl) Redirect(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

func NewStaticPageController() StaticPageController {
	return &StaticPageControllerImpl{}
}
