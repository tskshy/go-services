package routes

import (
	"github.com/gorilla/mux"
	"libs/log"
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
	log.Debug("test")
	var m = make(map[string]string)
	m["html"] = `"aa"<a>hello</a>`
	response.JResult(w, m)
	return
}
