package controller

import (
	"log"
	"net/http"
	"strconv"
	"task/helper"
	"task/model/web"
	"task/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("id")
	log.Println("product ID " + productId)
	id, err := strconv.Atoi(productId)
	helper.PanicHandling(err)
	response := controller.CategoryService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	queryValues := request.URL.Query()
	name := queryValues.Get("name")
	categories := controller.CategoryService.FindAll(request.Context(), name)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categories,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
