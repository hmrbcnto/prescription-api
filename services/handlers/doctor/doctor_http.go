package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/services/usecases"
)

type DoctorHttpHandler interface {
	InitRoutes(mux *mux.Router)
}

type doctor_http_handler struct {
	doctorUsecase usecases.DoctorUsecase
}

func NewDoctorHandler(doctorUsecase usecases.DoctorUsecase) DoctorHttpHandler {
	return &doctor_http_handler{
		doctorUsecase: doctorUsecase,
	}
}
