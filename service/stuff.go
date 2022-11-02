package service

import (
	"context"
	"fmt"
	"strconv"
	"task/exception"
	"task/helper"
	"task/model/domain"
	"task/model/web"
	"task/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type StuffService interface {
	Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponse
	Update(ctx context.Context, request web.TaskUpdateRequest) web.TaskResponse
	FindById(ctx context.Context, task int) web.TaskResponse
	FindTaskByTaskId(ctx context.Context, id int) []web.TaskResponse
	FindAll(ctx context.Context) []web.TaskResponse
	Delete(ctx context.Context, id int)
}

type StuffRepositoryImpl struct {
	StuffRepository repository.StuffRepo
	Validate        *validator.Validate
}

func NewStuffService(stuffRepo repository.StuffRepo, validate *validator.Validate) StuffService {
	return &StuffRepositoryImpl{
		StuffRepository: stuffRepo,
		Validate:        validate,
	}
}

func (service *StuffRepositoryImpl) Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponse {
	fmt.Printf("Parent" + strconv.Itoa(request.Parent_Task_Id))
	err := service.Validate.Struct(request)
	helper.PanicHandling(err)
	stuffDomain := domain.Task{
		Title:          request.Title,
		Image:          request.Image,
		Description:    request.Description,
		Parent_Task_Id: request.Parent_Task_Id,
		Poin:           int(request.Poin),
	}

	stuff := service.StuffRepository.CreateStuff(ctx, stuffDomain)
	return helper.StuffResponse(stuff)
}

func (service *StuffRepositoryImpl) Update(ctx context.Context, request web.TaskUpdateRequest) web.TaskResponse {
	err := service.Validate.Struct(request)
	helper.PanicHandling(err)

	_, errNotFound := service.StuffRepository.FindById(ctx, request.Id)
	helper.PanicHandling(errNotFound)

	stuffDomain := domain.Task{
		Id:             request.Id,
		Image:          request.Image,
		Description:    request.Description,
		Title:          request.Title,
		Created_At:     time.Now(),
		Updated_at:     time.Now(),
		Parent_Task_Id: int(request.Parent_Task_Id),
		Poin:           request.Poin,
	}

	stuff := service.StuffRepository.Update(ctx, stuffDomain)
	return helper.StuffResponse(stuff)
}

func (service *StuffRepositoryImpl) FindById(ctx context.Context, task int) web.TaskResponse {
	stuff, err := service.StuffRepository.FindById(ctx, task)
	if err != nil {
		panic(err)
	}
	return helper.StuffResponse(stuff)
}

func (service *StuffRepositoryImpl) FindTaskByTaskId(ctx context.Context, id int) []web.TaskResponse {
	stuffChildren := service.StuffRepository.FindTaskByTaskId(ctx, id)
	var responses []web.TaskResponse
	for _, item := range stuffChildren {
		responseStuf := web.TaskResponse{
			Id:             item.Id,
			Title:          item.Title,
			Description:    item.Description,
			Image:          item.Image,
			Parent_Task_Id: int(item.Parent_Task_Id),
			Updated_at:     item.Updated_at,
			Created_at:     item.Created_At,
			Poin:           item.Poin,
		}
		responses = append(responses, responseStuf)
	}
	return responses

}

func (service *StuffRepositoryImpl) FindAll(ctx context.Context) []web.TaskResponse {
	stuffs := service.StuffRepository.FindTask(ctx)
	var responses []web.TaskResponse
	for _, item := range stuffs {
		responseStuf := web.TaskResponse{
			Id:             item.Id,
			Title:          item.Title,
			Description:    item.Description,
			Image:          item.Image,
			Parent_Task_Id: int(item.Parent_Task_Id),
			Updated_at:     item.Updated_at,
			Created_at:     item.Created_At,
			Poin:           int(item.Poin),
		}

		responses = append(responses, responseStuf)
	}
	return responses
}

func (service *StuffRepositoryImpl) Delete(ctx context.Context, id int) {
	stuff, err := service.StuffRepository.FindById(ctx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.StuffRepository.Delete(ctx, stuff)
}
