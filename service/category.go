package service

import (
	"context"
	"task/helper"
	"task/model/web"
	"task/repository/category"

	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	FindAll(ctx context.Context, name string) []web.CategoryResponse
	FindById(ctx context.Context, id int) web.CategoryResponse
	Delete(ctx context.Context, id int)
}

type CategoryRepositoryImpl struct {
	CategoryRepository category.CategoryRepo
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepo category.CategoryRepo, validate *validator.Validate) CategoryService {
	return &CategoryRepositoryImpl{
		CategoryRepository: categoryRepo,
		Validate:           validate,
	}
}

func (service *CategoryRepositoryImpl) FindAll(ctx context.Context, name string) []web.CategoryResponse {
	categories := service.CategoryRepository.FindByAll(ctx, name)
	var responses []web.CategoryResponse
	for _, item := range categories {
		var responseProducts []web.ProductResponse
		for _, v := range item.Products {
			responseProduct := web.ProductResponse{
				Id:    v.Id,
				Name:  v.Name,
				Stock: v.Stock,
				Price: v.Price,
			}
			responseProducts = append(responseProducts, responseProduct)
		}

		responseStuf := web.CategoryResponse{
			Id:         item.Id,
			Updated_at: item.Updated_at,
			Created_at: item.Created_At,
			Name:       item.Name,
			Products:   responseProducts,
		}

		responses = append(responses, responseStuf)
	}
	return responses
}

func (service *CategoryRepositoryImpl) FindById(ctx context.Context, id int) web.CategoryResponse {
	category, err := service.CategoryRepository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}
	return helper.CategoryResponse(category)
}

func (service *CategoryRepositoryImpl) Delete(ctx context.Context, id int) {
	panic("not implemented") // TODO: Implement
}
