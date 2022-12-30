package app

import (
	"task/controller"
	"task/exception"

	"github.com/julienschmidt/httprouter"
	group "github.com/mythrnr/httprouter-group"
)

func NewRouter(stuffController controller.StuffController,
	auth controller.AuthUserHandler,
	productController controller.ProductController,
	categoryController controller.CategoryController,
) *httprouter.Router {
	router := httprouter.New()

	group.New("/api").Middleware()

	router.GET("/api/tasks", stuffController.FindAll)
	router.GET("/api/carts/:id", stuffController.FindById)
	router.POST("/api/carts", stuffController.Create)
	router.PUT("/api/carts/:id", stuffController.Update)
	router.DELETE("/api/carts/:id", stuffController.Delete)
	router.POST("/api/auth/register", auth.Create)
	router.POST("/api/auth/login", auth.Login)

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:id", categoryController.FindById)

	router.GET("/api/product/:id", productController.FindById)
	router.GET("/api/products", productController.FindByName)
	router.PanicHandler = exception.ErrorHandler
	return router
}
