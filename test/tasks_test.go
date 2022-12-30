// package test

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"task/app"
// 	"task/controller"
// 	"task/exception"
// 	"task/helper"
// 	"task/model/domain"
// 	"task/repository"
// 	"task/service"
// 	"testing"

// 	"github.com/go-playground/assert/v2"
// 	"github.com/go-playground/validator/v10"
// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type ConfigDB struct {
// 	DB_Username string
// 	DB_Password string
// 	DB_Host     string
// 	DB_Port     string
// 	DB_Database string
// }

// func ConnectDB(config *ConfigDB) *gorm.DB {
// 	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
// 		config.DB_Host,
// 		config.DB_Username,
// 		config.DB_Password,
// 		config.DB_Database,
// 		config.DB_Port,
// 	)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	helper.PanicHandling(err)
// 	return db
// }

// func Configuration() *ConfigDB {
// 	error := godotenv.Load("../.env")
// 	helper.PanicHandling(error)
// 	configDB := ConfigDB{
// 		DB_Username: os.Getenv("DB_USERNAME"),
// 		DB_Password: os.Getenv("DB_PASSWORD"),
// 		DB_Host:     os.Getenv("DB_HOST"),
// 		DB_Port:     os.Getenv("DB_PORT"),
// 		DB_Database: os.Getenv("DB_NAME"),
// 	}
// 	return &configDB
// }

// func Setup() http.Handler {

// 	validate := validator.New()
// 	db := ConnectDB(Configuration())
// 	repo := repository.NewStuffRepository(db)
// 	service := service.NewStuffService(repo, validate)
// 	controller := controller.NewStuffController(service)
// 	router := app.NewRouter(controller)
// 	router.PanicHandler = exception.ErrorHandler
// 	return router
// }

// func Truncate(db gorm.DB) {
// 	db.Exec("TRUNCATE ONLY tasks RESTART IDENTITY")
// }

// func Test_GetByIDSuccess(t *testing.T) {
// 	db := ConnectDB(Configuration())
// 	Truncate(*db)

// 	tx := db.Begin()
// 	stuffRepository := repository.NewStuffRepository(db)
// 	stuff := stuffRepository.CreateStuff(context.Background(), domain.Task{
// 		Title:       "task one",
// 		Description: "description task one",
// 		Image:       "image.jpg",
// 		Poin:        3,
// 	})
// 	tx.Commit()

// 	router := Setup()
// 	route := "http://" + os.Getenv("APP_URL") + "/api/tasks/" + strconv.Itoa(stuff.Id)
// 	request := httptest.NewRequest(http.MethodGet, route, nil)
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, request)
// 	response := recorder.Result()

// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "Success", responseBody["status"])

// 	assert.Equal(t, stuff.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
// 	assert.Equal(t, stuff.Title, responseBody["data"].(map[string]interface{})["title"])
// 	assert.Equal(t, stuff.Description, responseBody["data"].(map[string]interface{})["description"])
// 	assert.Equal(t, stuff.Image, responseBody["data"].(map[string]interface{})["image"])
// }

// func Test_Delete(t *testing.T) {
// 	db := ConnectDB(Configuration())
// 	Truncate(*db)

// 	tx := db.Begin()
// 	stuffRepository := repository.NewStuffRepository(db)
// 	stuff := stuffRepository.CreateStuff(context.Background(), domain.Task{
// 		Title:       "task one",
// 		Description: "description task one",
// 		Image:       "image.jpg",
// 		Poin:        3,
// 	})
// 	tx.Commit()

// 	router := Setup()
// 	route := "http://" + os.Getenv("APP_URL") + "/api/tasks/" + strconv.Itoa(stuff.Id)
// 	request := httptest.NewRequest(http.MethodDelete, route, nil)
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, request)
// 	response := recorder.Result()

// 	assert.Equal(t, 200, response.StatusCode)
// }

// func Test_Create(t *testing.T) {
// 	router := Setup()
// 	requestBody := strings.NewReader(`{
// 		"title" : "stuff one",
// 		"description" : "stuff one",
// 		"image":"lorem_ipsum.jpg",
// 		"parent_task_id": null 	,
// 		"poin": 2
// 	}`)

// 	route := "http://" + os.Getenv("APP_URL") + "/api/tasks"
// 	request := httptest.NewRequest(http.MethodPost, route, requestBody)
// 	request.Header.Add("Content-Type", "application/json")

// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, request)
// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	fmt.Println(responseBody)

// 	assert.Equal(t, "Success", responseBody["status"])
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// }
