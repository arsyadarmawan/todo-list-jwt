package domain

import (
	"time"
)

type Category struct {
	Id         int
	Name       string
	Created_At time.Time
	Updated_at time.Time
	Products   []Product `json:"products" gorm:"many2many:categories_products"`
}
