package web

import "time"

type TaskResponse struct {
	Id         int       `json:"id"`
	User_Id    int       `json:"user_id"`
	Product_Id int       `json:"product_id"`
	Total      int       `json:"total"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type ProductResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type CategoryResponse struct {
	Id         int               `json:"id"`
	Name       string            `json:"title"`
	Created_at time.Time         `json:"created_at"`
	Updated_at time.Time         `json:"updated_at"`
	Products   []ProductResponse `json:"products" gorm:"many2many:categories_products"`
}
