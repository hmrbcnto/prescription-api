package router

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/auth_repo"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/doctor_repo"
	auth_handler "github.com/hmrbcnto/prescription-api/services/handlers/auth"
	doctor_handler "github.com/hmrbcnto/prescription-api/services/handlers/doctor"
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
	userHttpHandler := doctor_handler.NewDoctorHandler(doctorUsecase)

	authRepo := auth_repo.NewRepo(db)
	authUsecase := usecases.NewAuthUsecase(authRepo)
	authHttpHandler := auth_handler.NewAuthHandler(authUsecase)

	routers := []router{userHttpHandler, authHttpHandler}

	for _, router := range routers {
		router.InitRoutes(mux)
	}
}
