package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"task/helper"
	"task/model/web"
	"task/service"

	"task/middleware"

	"github.com/julienschmidt/httprouter"
)

type StuffController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type StuffControllerImpl struct {
	StuffService service.StuffService
	AuthService  service.AuthService
}

func NewStuffController(service service.StuffService, autService service.AuthService) StuffController {
	return &StuffControllerImpl{
		StuffService: service,
		AuthService:  autService,
	}
}

func (controller *StuffControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authToken := request.Header.Get("Authorization")
	claims, isTrue := middleware.ExtractClaims(authToken)

	username := claims["username"].(string)
	user := controller.AuthService.CheckUsername(request.Context(), username)

	if authToken == "" || isTrue == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	var msg web.TaskCreateRequest
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	msg.User_Id = user.Id

	response := controller.StuffService.Create(request.Context(), msg)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StuffControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authToken := request.Header.Get("Authorization")
	_, isTrue := middleware.ExtractClaims(authToken)
	if authToken == "" || isTrue == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	requestStuff := web.TaskUpdateRequest{}
	helper.ReadFromRequestBody(request, &requestStuff)
	stuffId := params.ByName("id")
	id, err := strconv.Atoi(stuffId)
	helper.PanicHandling(err)

	requestStuff.Id = id

	response := controller.StuffService.Update(request.Context(), requestStuff)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StuffControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authToken := request.Header.Get("Authorization")
	_, isTrue := middleware.ExtractClaims(authToken)
	if authToken == "" || isTrue == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	stuffId := params.ByName("id")
	id, err := strconv.Atoi(stuffId)
	helper.PanicHandling(err)

	controller.StuffService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   "Row deleted",
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *StuffControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authToken := request.Header.Get("Authorization")
	_, isTrue := middleware.ExtractClaims(authToken)
	if authToken == "" || isTrue == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	stuffId := params.ByName("id")
	id, err := strconv.Atoi(stuffId)
	helper.PanicHandling(err)
	response := controller.StuffService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StuffControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authToken := request.Header.Get("Authorization")
	claims, isTrue := middleware.ExtractClaims(authToken)
	if authToken == "" || isTrue == false {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Invalidate Token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	username := claims["username"].(string)
	user := controller.AuthService.CheckUsername(request.Context(), username)

	responses := controller.StuffService.FindAll(request.Context(), user.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   responses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
