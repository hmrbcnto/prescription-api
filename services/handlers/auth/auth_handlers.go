package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/models"
)

func (authHandler *auth_http_handler) login(w http.ResponseWriter, r *http.Request) {
	// Get request body
	loginData := new(models.LoginStruct)
	err := json.NewDecoder(r.Body).Decode(loginData)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)

	loginResults, err := authHandler.authUsecase.Login(loginData.Username, loginData.Password)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: loginResults.TokenString,
		Expires: loginResults.ExpirationTime,
	})

	jsonWriter.Encode(loginResults)
}

func (authHandler *auth_http_handler) InitRoutes(mux *mux.Router) {
	subRouter := mux.PathPrefix("").Subrouter()

	// Generate routes
	subRouter.HandleFunc("/login", authHandler.login).Methods("POST")
}
