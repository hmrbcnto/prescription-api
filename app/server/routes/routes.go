package router

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/auth_repo"
	"github.com/hmrbcnto/prescription-api/infastructure/db/mongo/user_repo"
	auth_handler "github.com/hmrbcnto/prescription-api/services/handlers/auth"
	user_handler "github.com/hmrbcnto/prescription-api/services/handlers/user"
	"github.com/hmrbcnto/prescription-api/services/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type router interface {
	InitRoutes(mux *mux.Router)
}

func InitializeRoutes(db *mongo.Client, mux *mux.Router) {
	// Generate http handlers
	userRepo := user_repo.NewRepo(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHttpHandler := user_handler.NewUserHandler(userUsecase)

	authRepo := auth_repo.NewRepo(db)
	authUsecase := usecases.NewAuthUsecase(authRepo)
	authHttpHandler := auth_handler.NewAuthHandler(authUsecase)

	routers := []router{userHttpHandler, authHttpHandler}

	for _, router := range routers {
		router.InitRoutes(mux)
	}
}
