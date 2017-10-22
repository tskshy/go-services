package routes

import (
	"github.com/gorilla/mux"
	"libs3/response"
	"net/http"
)

func Init() *mux.Router {
	var router = mux.NewRouter()
	routes(router)
	return router
}

func routes(router *mux.Router) {
	router.HandleFunc("/", response.DecoratorFunc(home)).Methods("GET")
}

var home = func(w http.ResponseWriter, r *http.Request) {
	panic("error 手动")
	response.JOK(w)
	return
}
