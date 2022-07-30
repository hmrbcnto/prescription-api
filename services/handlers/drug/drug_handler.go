package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/entities"
)

func (drugHandler *drug_http_handler) getAllDrugs(w http.ResponseWriter, r *http.Request) {
	// Get all drugs
	w.Header().Set("Content-Type", "application/json")
	drugs, err := drugHandler.drugUsecase.GetAllDrugs()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter := json.NewEncoder(w)
	jsonWriter.Encode(drugs)
}

func (drugHandler *drug_http_handler) getDrugById(w http.ResponseWriter, r *http.Request) {
	// Setting header
	w.Header().Set("Content-Type", "application/json")

	// Get params
	params := mux.Vars(r)
	id := params["id"]

	user, err := drugHandler.drugUsecase.GetDrugById(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Encoding
	jsonWriter := json.NewEncoder(w)
	jsonWriter.Encode(user)
}

func (drugHandler *drug_http_handler) createDrug(w http.ResponseWriter, r *http.Request) {
	// Setting header
	w.Header().Set("Content-Type", "application/json")

	// Getting request body
	drug := new(entities.Drug)
	err := json.NewDecoder(r.Body).Decode(drug)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	createdDrug, err := drugHandler.drugUsecase.CreateDrug(drug)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Encoding to json
	jsonWriter := json.NewEncoder(w)
	jsonWriter.Encode(createdDrug)
}

func (drugHandler *drug_http_handler) deleteDrugById(w http.ResponseWriter, r *http.Request) {
	// Setting header
	w.Header().Set("Content-Type", "application/json")
	
	// Getting params
	params := mux.Vars(r)
	id := params["id"]

	err := drugHandler.drugUsecase.DeleteDrugById(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}

func (drugHandler *drug_http_handler) updateDrugById(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Getting params
	params := mux.Vars(r)
	id := params["id"]

	// Getting body
	drug := new(entities.Drug)
	err := json.NewDecoder(r.Body).Decode(drug)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Updating
	updatedDrug, err := drugHandler.drugUsecase.UpdateDrugById(id, drug)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Encoding
	jsonWriter := json.NewEncoder(w)
	jsonWriter.Encode(updatedDrug)
}

func (drugHandler *drug_http_handler) InitRoutes(mux *mux.Router) {
	subRouter := mux.PathPrefix("/drugs").Subrouter()

	// Generate routes
	subRouter.HandleFunc("", drugHandler.getAllDrugs).Methods("GET")
	subRouter.HandleFunc("/id", drugHandler.getDrugById).Methods("GET")
	subRouter.HandleFunc("", drugHandler.createDrug).Methods("POST")
	subRouter.HandleFunc("/id", drugHandler.updateDrugById).Methods("PUT")
	subRouter.HandleFunc("/id", drugHandler.deleteDrugById).Methods("DELETE")
}