package routes

import (
	"net/http"
	"npm-repo/controllers"
	"npm-repo/services"

	"github.com/gorilla/mux"
)

// Route is
type Route struct{}

func (r *Route) Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	npmService := new(services.NPMService)
	npmController := controllers.InitNPMController(npmService)

	router.HandleFunc("/read/{npm}", npmController.GETNPM).Methods(http.MethodGet)
	router.HandleFunc("/read/{npm}/{trx-id}", npmController.GETNPM).Methods(http.MethodGet)

	return router
}
