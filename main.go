package main

import (
	"fmt"
	"net/http"
	"os"
	"task/app"
	"task/controller"
	"task/exception"
	"task/helper"
	"task/repository"
	"task/repository/category"
	"task/repository/product"

	"task/service"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	validate := validator.New()
	error := godotenv.Load(".env")
	helper.PanicHandling(error)

	configDB := app.ConfigDB{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Database: os.Getenv("DB_NAME"),
	}

	db := app.NewDB(&configDB)

	// users
	repoUser := repository.NewAuthRepo(db)
	serviceUser := service.NewUserService(repoUser, validate)
	controllerUser := controller.NewAuthController(serviceUser)
	// tasks
	repo := repository.NewStuffRepository(db)
	services := service.NewStuffService(repo, validate)
	controllers := controller.NewStuffController(services, serviceUser)

	// product
	repoProduct := product.NewProductRepository(db)
	productService := service.NewProductService(repoProduct, validate)
	controllerProduct := controller.NewProductController(productService)

	// categories
	categoryRepo := category.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(controllers, &controllerUser, controllerProduct, categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    os.Getenv("APP_URL"),
		Handler: router,
	}
	fmt.Printf("Running on server " + os.Getenv("APP_URL") + "\n")
	err := server.ListenAndServe()
	helper.PanicHandling(err)

}
