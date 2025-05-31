package app

import (
	"goweb/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router(authController controller.AuthController, todoController controller.TodoController, staticController controller.StaticPageController) *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(staticController.NotFound)

	router.GET("/", staticController.IndexPage)

	router.POST("/api/login", authController.Login)
	router.POST("/api/logout", authController.Logout)
	router.POST("/api/register", authController.RegisterUser)

	router.GET("/api/todo/list", todoController.ListTodo)
	router.POST("/api/todo/create", todoController.CreateTodo)
	router.POST("/api/todo/modify/:id", todoController.ModifyTodo)
	router.GET("/api/todo/get/:id", todoController.SingleTodo)
	router.DELETE("/api/todo/delete/:id", todoController.DeleteTodo)
	router.GET("/api/todo/all", todoController.AllTodo)

	return router
}

// func AuthRouter(authController controller.AuthController) http.Handler {

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
// 		writer.Write([]byte("Yeayy index!"))
// 	})

// 	mux.HandleFunc("/api/login", authController.Login)

// 	mux.HandleFunc("/api/logout", authController.Logout)

// 	return mux
// }
