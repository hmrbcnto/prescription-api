package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/services/usecases"
)

type DrugHttpHandler interface {
	InitRoutes(mux *mux.Router)
}

type drug_http_handler struct {
	drugUsecase usecases.DrugUsecase
}

func NewDrugHandler(drugUsecase usecases.DrugUsecase) DrugHttpHandler {
	return &drug_http_handler{
		drugUsecase: drugUsecase,
	}
}