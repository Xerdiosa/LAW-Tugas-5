package controllers

import (
	"encoding/json"
	"net/http"
	"npm-repo/helpers"
	"npm-repo/services"

	"github.com/gorilla/mux"
)

func InitNPMController(npmService services.INPMService) *NPMController {

	npmController := new(NPMController)
	npmController.NPMService = npmService

	return npmController
}

type NPMController struct {
	NPMService services.INPMService
}

func (n *NPMController) GETNPM(res http.ResponseWriter, req *http.Request) {
	npm := mux.Vars(req)["npm"]
	name, err := n.NPMService.GetNPM(npm)
	if err != nil {
		response := helpers.ReadResponse{
			Status: "ERROR",
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}

	response := helpers.ReadResponse{
		Status: "OK",
		Name:   name,
		NPM:    npm,
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)
}
