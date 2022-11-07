package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type Token interface {
	GenerateToken(username string) (string, string)
	ValidateToken(token string) (*jwt.Token, error)
}

type token struct {
	secretKey string
}

type authClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewToken(secretKey string) Token {
	return &token{secretKey: secretKey}
}

func (t *token) GenerateToken(username string) (string, string) {
	claims := &authClaims{
		username,
		jwt.StandardClaims{},
	}

	ctx := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := ctx.SignedString([]byte(t.secretKey))
	if err != nil {
		logrus.Panic(err)
	}

	return time.Now().Add(time.Hour * 2).Format(time.RFC3339), token
}

func (t *token) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}
		return []byte(t.secretKey), nil
	})
}
