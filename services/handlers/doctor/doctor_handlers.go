package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/entities"
	"github.com/hmrbcnto/prescription-api/middlewares"
)

func (doctorHandler *doctor_http_handler) createDoctor(w http.ResponseWriter, r *http.Request) {
	// Get http request body
	doctor := new(entities.Doctor)
	err := json.NewDecoder(r.Body).Decode(doctor)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)

	createdUser, err := doctorHandler.doctorUsecase.CreateDoctor(doctor)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter.Encode(createdUser)
}

func (doctorHandler *doctor_http_handler) getAllDoctors(w http.ResponseWriter, r *http.Request) {
	// Get all users
	w.Header().Set("Content-Type", "application/json")
	users, err := doctorHandler.doctorUsecase.GetAllDoctors()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter := json.NewEncoder(w)
	jsonWriter.Encode(users)
}

func (doctorHandler *doctor_http_handler) getDoctorById(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)

	// Get params
	params := mux.Vars(r)
	userId := params["id"]

	user, err := doctorHandler.doctorUsecase.GetDoctorById(userId)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter.Encode(user)
}

func (doctorHandler *doctor_http_handler) InitRoutes(mux *mux.Router) {
	subRouter := mux.PathPrefix("/doctors").Subrouter()

	// Generate routes
	subRouter.Use(middlewares.CheckForToken)
	subRouter.HandleFunc("", doctorHandler.getAllDoctors).Methods("GET")
	subRouter.HandleFunc("/id", doctorHandler.getDoctorById).Methods("GET")
	subRouter.HandleFunc("", doctorHandler.createDoctor).Methods("POST")
}
