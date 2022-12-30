package service

import (
	"context"
	"task/helper"
	"task/model/web"
	"task/repository/product"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	FindByName(ctx context.Context, name string) []web.ProductResponse
	FindById(ctx context.Context, task int) web.ProductResponse
	Delete(ctx context.Context, id int)
}

func NewProductService(productRepo product.ProductRepo, validate *validator.Validate) ProductService {
	return &ProductRepositoryImpl{
		ProductRepository: productRepo,
		Validate:          validate,
	}
}

type ProductRepositoryImpl struct {
	ProductRepository product.ProductRepo
	Validate          *validator.Validate
}

func (service *ProductRepositoryImpl) FindByName(ctx context.Context, name string) []web.ProductResponse {
	products := service.ProductRepository.FindByName(ctx, name)
	var responses []web.ProductResponse
	for _, v := range products {
		product := web.ProductResponse{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Created_at:  v.Created_At,
			Updated_at:  v.Updated_at,
			Stock:       v.Stock,
			Price:       v.Price,
		}
		responses = append(responses, product)
	}
	return responses
}

func (service *ProductRepositoryImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	product, err := service.ProductRepository.FindById(ctx, productId)
	if err != nil {
		panic(err)
	}
	return helper.ProductResponse(product)
}

func (service *ProductRepositoryImpl) Delete(ctx context.Context, id int) {
	panic("not implemented") // TODO: Implement
}
