package controller

import (
	"Marcketplace/model/entities"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(user entities.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	return token
}
