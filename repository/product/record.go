package product

import (
	_domain "task/model/domain"
	"time"
)

type Product struct {
	Id          uint `gorm:"primaryKey autoIncrement"`
	Name        string
	Description string
	Price       float64
	Stock       int
	Category_id int
	Created_At  time.Time
	Updated_at  time.Time
}

func (p *Product) ToDomain() _domain.Product {
	return _domain.Product{
		Id:          int(p.Id),
		Name:        p.Name,
		Description: p.Description,
		Category_id: int(p.Category_id),
		Stock:       int(p.Stock),
		Price:       float64(p.Price),
		Created_At:  p.Created_At,
		Updated_at:  p.Updated_at,
	}
}
