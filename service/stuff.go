package service

import (
	"context"
	"task/exception"
	"task/helper"
	"task/model/domain"
	"task/model/web"
	"task/repository"

	"github.com/go-playground/validator/v10"
)

type StuffService interface {
	Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponse
	Update(ctx context.Context, request web.TaskUpdateRequest) web.TaskResponse
	FindById(ctx context.Context, task int) web.TaskResponse
	FindAll(ctx context.Context, userId int) []web.TaskResponse
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
	err := service.Validate.Struct(request)
	helper.PanicHandling(err)
	stuffDomain := domain.Task{
		User_id:    request.User_Id,
		Product_id: request.Product_Id,
		Total:      request.Total,
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
		Id:    request.Id,
		Total: request.Total,
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

func (service *StuffRepositoryImpl) FindAll(ctx context.Context, userId int) []web.TaskResponse {
	stuffs := service.StuffRepository.FindTask(ctx, userId)
	var responses []web.TaskResponse
	for _, item := range stuffs {
		responseStuf := web.TaskResponse{
			Id:         item.Id,
			User_Id:    item.User_id,
			Product_Id: item.Product_id,
			Updated_at: item.Updated_at,
			Created_at: item.Created_At,
			Total:      int(item.Total),
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
