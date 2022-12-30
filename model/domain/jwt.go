package domain

import (
	_response "task/model/web"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Id       int    `json:id`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenJWT struct {
	AccesToken string
	User       _response.UserReponse
}
