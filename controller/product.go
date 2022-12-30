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

type ProductController interface {
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: service,
	}
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("id")
	id, err := strconv.Atoi(productId)
	helper.PanicHandling(err)
	log.Println("product ID " + productId)
	response := controller.ProductService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	queryValues := request.URL.Query()
	name := queryValues.Get("name")
	response := controller.ProductService.FindByName(request.Context(), name)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
