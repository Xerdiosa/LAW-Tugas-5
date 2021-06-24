package controllers

import (
	"encoding/json"
	"net/http"
	"npm-repo/helpers"
	"npm-repo/services"
)

func InitNPMController(npmService services.INPMService) *NPMController {

	npmController := new(NPMController)
	npmController.NPMService = npmService

	return npmController
}

type NPMController struct {
	NPMService services.INPMService
}

func (n *NPMController) ADDNPM(res http.ResponseWriter, req *http.Request) {
	var dataBody helpers.UpdateRequest
	err := json.NewDecoder(req.Body).Decode(&dataBody)
	if err != nil {
		response := helpers.UpdateResponse{
			Status: "ERROR",
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}

	err = n.NPMService.AddNPM(dataBody.NPM, dataBody.Name)
	if err != nil {
		response := helpers.UpdateResponse{
			Status: "ERROR",
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}

	response := helpers.UpdateResponse{
		Status: "OK",
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)
}
