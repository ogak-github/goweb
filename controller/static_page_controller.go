package controller

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StaticPageController interface {
	IndexPage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	NotFound(writer http.ResponseWriter, request *http.Request)
	Redirect(writer http.ResponseWriter, request *http.Request)
}

type StaticPageControllerImpl struct {
	staticContent embed.FS
}

// IndexPage implements StaticPageController.
func (s *StaticPageControllerImpl) IndexPage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	///// THIS NOT WORKING ON GO BINARY/BUILD
	//http.ServeFile(writer, request, "views/index.html")

	// Split static files (css/js/images) dari embed FS
	staticRoot, err := fs.Sub(s.staticContent, "static")
	if err != nil {
		log.Fatal(err)
	}

	// Serve file static (for example router: /static/*filepath)
	staticServer := http.StripPrefix("/static/", http.FileServer(http.FS(staticRoot)))
	http.Handle("/static/", staticServer)

	// Load index.html from embed FS
	htmlRoot, err := fs.Sub(s.staticContent, "views")
	if err != nil {
		log.Fatal(err)
	}

	// Open file index.html
	indexFile, err := htmlRoot.Open("index.html")
	if err != nil {
		http.Error(writer, "Index page not found", http.StatusInternalServerError)
		return
	}
	defer indexFile.Close()

	// Serve file HTML
	http.ServeFileFS(writer, request, htmlRoot, "index.html")

}

// NotFound implements StaticPageController.
func (s *StaticPageControllerImpl) NotFound(writer http.ResponseWriter, request *http.Request) {
	// Split static files (css/js/images) from embed FS
	staticRoot, err := fs.Sub(s.staticContent, "static")
	if err != nil {
		log.Fatal(err)
	}

	// Serve file static (for example router: /static/*filepath)
	staticServer := http.StripPrefix("/static/", http.FileServer(http.FS(staticRoot)))
	http.Handle("/static/", staticServer)

	// Load index.html from embed FS
	htmlRoot, err := fs.Sub(s.staticContent, "views")
	if err != nil {
		log.Fatal(err)
	}

	// Open file index.html
	indexFile, err := htmlRoot.Open("404.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(writer, "404 page not found", http.StatusInternalServerError)
		return
	}
	defer indexFile.Close()

	// Serve file HTML
	http.ServeFileFS(writer, request, htmlRoot, "404.html")
}

// Redirect implements StaticPageController.
func (s *StaticPageControllerImpl) Redirect(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

func NewStaticPageController(staticContent embed.FS) StaticPageController {
	return &StaticPageControllerImpl{
		staticContent: staticContent,
	}
}
