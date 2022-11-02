package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"task/helper"
	"task/model/web"
	"task/service"

	"github.com/julienschmidt/httprouter"
)

type StuffController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindSubTaskBYID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type StuffControllerImpl struct {
	StuffService service.StuffService
}

func NewStuffController(service service.StuffService) StuffController {
	return &StuffControllerImpl{
		StuffService: service,
	}
}

func (controller *StuffControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createRequest := web.TaskCreateRequest{}
	helper.ReadFromRequestBody(request, &createRequest)

	response := controller.StuffService.Create(request.Context(), createRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StuffControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestStuff := web.TaskUpdateRequest{}
	fmt.Printf(requestStuff.Title)
	helper.ReadFromRequestBody(request, &requestStuff)
	fmt.Printf("halloo")
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
	responses := controller.StuffService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   responses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StuffControllerImpl) FindSubTaskBYID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stuffId := params.ByName("id")
	id, err := strconv.Atoi(stuffId)
	helper.PanicHandling(err)

	response := controller.StuffService.FindTaskByTaskId(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
