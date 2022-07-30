package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hmrbcnto/prescription-api/entities"
)

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginReturn struct {
	User entities.User
	TokenString string
	ExpirationTime time.Time
}