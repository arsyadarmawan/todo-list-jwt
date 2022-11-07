package app

import (
	"task/controller"
	"task/exception"

	"github.com/julienschmidt/httprouter"
	group "github.com/mythrnr/httprouter-group"
)

func NewRouter(stuffController controller.StuffController, auth controller.AuthUserHandler) *httprouter.Router {
	router := httprouter.New()

	group.New("/api").Middleware()

	router.GET("/api/tasks", stuffController.FindAll)
	router.GET("/api/tasks/:id", stuffController.FindById)
	router.GET("/api/subtaks/:id", stuffController.FindSubTaskBYID)
	router.POST("/api/tasks", stuffController.Create)
	router.PUT("/api/tasks/:id", stuffController.Update)
	router.DELETE("/api/tasks/:id", stuffController.Delete)
	router.POST("/api/auth/register", auth.Create)
	router.POST("/api/auth/login", auth.Login)

	router.PanicHandler = exception.ErrorHandler
	return router
}
