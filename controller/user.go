package controller

import (
	"fmt"
	"net/http"
	"os"
	"task/helper"
	"task/model/web"
	"task/service"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Id       int    `json:id`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthUserHandler interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) UserController {
	return UserController{
		authService: service,
	}
}

func (a *UserController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &createRequest)

	response := a.authService.Register(request.Context(), createRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (a *UserController) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	createRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &createRequest)

	hashPass, err := a.authService.Login(request.Context(), createRequest)
	helper.PanicHandling(err)
	fmt.Printf("BANGSAAATTT" + createRequest.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(createRequest.Password)); err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Failed to Login",
			Data:   "User or Password doesnt match",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	userDomain := a.authService.CheckUsername(request.Context(), createRequest.Username)

	claims := &jwtCustomClaims{
		userDomain.Id,
		userDomain.Name,
		userDomain.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	helper.PanicHandling(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Found",
		Data:   t,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (a *UserController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
