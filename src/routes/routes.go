package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"api/response"
)

func Init() *mux.Router {
	var router = mux.NewRouter()
	routes(router)
	return router
}

func routes(router *mux.Router) {
	router.HandleFunc("/", response.DecoratorFunc(home)).Methods("GET")
	router.HandleFunc("/", response.DecoratorFunc(home)).Methods("POST")
}

var home = func(w http.ResponseWriter, r *http.Request) {
	response.JOK(w)
	return
}
