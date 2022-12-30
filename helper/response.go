package helper

import (
	"task/model/domain"
	"task/model/web"
)

func StuffResponse(stuff domain.Task) web.TaskResponse {
	return web.TaskResponse{
		Id:         stuff.Id,
		Created_at: stuff.Created_At,
		Updated_at: stuff.Updated_at,
		Total:      stuff.Total,
		User_Id:    stuff.User_id,
		Product_Id: stuff.Product_id,
	}
}

func ProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.Id,
		Description: product.Description,
		Created_at:  product.Created_At,
		Updated_at:  product.Updated_at,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func CategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:         category.Id,
		Updated_at: category.Updated_at,
		Name:       category.Name,
		Products:   []web.ProductResponse{},
	}
}
