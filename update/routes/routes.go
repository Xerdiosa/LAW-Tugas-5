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

	router.HandleFunc("/update", npmController.ADDNPM).Methods(http.MethodPost)

	return router
}
