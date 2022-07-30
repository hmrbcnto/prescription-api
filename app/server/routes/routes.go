package router

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/auth_repo"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/doctor_repo"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/drug_repo"
	auth_handler "github.com/hmrbcnto/prescription-api/services/handlers/auth"
	doctor_handler "github.com/hmrbcnto/prescription-api/services/handlers/doctor"
	handlers "github.com/hmrbcnto/prescription-api/services/handlers/drug"
	"github.com/hmrbcnto/prescription-api/services/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type router interface {
	InitRoutes(mux *mux.Router)
}

func InitializeRoutes(db *mongo.Client, mux *mux.Router) {
	// Generate http handlers
	doctorRepo := doctor_repo.NewRepo(db)
	doctorUsecase := usecases.NewDoctorUsecase(doctorRepo)
	doctorHttpHandler := doctor_handler.NewDoctorHandler(doctorUsecase)

	drugRepo := drug_repo.NewRepo(db)
	drugUsecase := usecases.NewDrugUsecase(drugRepo)
	drugHttpHandler := handlers.NewDrugHandler(drugUsecase)

	authRepo := auth_repo.NewRepo(db)
	authUsecase := usecases.NewAuthUsecase(authRepo)
	authHttpHandler := auth_handler.NewAuthHandler(authUsecase)

	routers := []router{doctorHttpHandler, authHttpHandler, drugHttpHandler}

	for _, router := range routers {
		router.InitRoutes(mux)
	}
}
