package app

import (
	"task/controller"
	"task/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(stuffController controller.StuffController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tasks", stuffController.FindAll)
	router.GET("/api/tasks/:id", stuffController.FindById)
	router.GET("/api/subtaks/:id", stuffController.FindSubTaskBYID)
	router.POST("/api/tasks", stuffController.Create)
	router.PUT("/api/tasks/:id", stuffController.Update)
	router.DELETE("/api/tasks/:id", stuffController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
