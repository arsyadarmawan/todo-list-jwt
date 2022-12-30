package domain

import "time"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Stock       int
	Category_id int
	Created_At  time.Time
	Updated_at  time.Time
	Categories  []*Category `json:"categories" gorm:"many2many:categories_products"`
}

type CategoryProduct struct {
	Category_id int
	Product_id  int
	Created_At  time.Time
	Updated_at  time.Time
}
