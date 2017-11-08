package routes

import (
	"libs3/response"
	"net/http"

	"github.com/gorilla/mux"

	"app/conf"
	"app/services/home"
)

func Init() *mux.Router {
	var router = mux.NewRouter().StrictSlash(true)
	routes(router)
	return router
}

/*add routes info*/
func routes(router *mux.Router) {
	router.NotFoundHandler = &response.NotFound
	router.MethodNotAllowedHandler = &response.MethodNotAllowed

	/*静态资源*/
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(conf.AppConfInfo.WWWDir)))).Methods("GET")
	/*首页*/
	router.HandleFunc("/", response.DecoratorFunc(home.HomePage)).Methods("GET")
}
