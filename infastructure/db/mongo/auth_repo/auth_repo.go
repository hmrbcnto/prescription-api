package auth_repo

import (
	"context"
	"errors"
	"time"

	"github.com/hmrbcnto/prescription-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo interface {
	Login(username string, password string) (*entities.Doctor, error)
}

type authRepo struct {
	db *mongo.Collection
}

func NewRepo(db *mongo.Client) AuthRepo {
	return &authRepo{
		db: db.Database("prescription-api").Collection("doctors"),
	}
}

func (ar *authRepo) Login(username string, password string) (*entities.Doctor, error) {

	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating query
	// query := bson.M{"username": username}
	query := bson.D{{Key: "username", Value: username}}

	// Querying
	result := ar.db.FindOne(ctx, query)

	// Unmarshal/Decode into user object
	foundDoctor := new(entities.Doctor)
	err := result.Decode(foundDoctor)

	if err != nil {
		return nil, err
	}

	// Check if user's password matches given password
	if foundDoctor.Password != password {
		return nil, errors.New("invalid password")
	}

	// Returning
	return foundDoctor, nil
}
