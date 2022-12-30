package category

import (
	_domain "task/model/domain"
	"task/repository/product"
	"time"
)

type Category struct {
	Id         int
	Name       string
	Created_At time.Time
	Updated_at time.Time
	Products   []product.Product `json:"products" gorm:"many2many:categories_products"`
}

func (c *Category) ToDomain() _domain.Category {
	return _domain.Category{
		Id:         int(c.Id),
		Name:       c.Name,
		Created_At: c.Created_At,
		Updated_at: c.Updated_at,
		Products:   c.ToDomainProduct(),
	}
}

func (c *Category) ToDomainProduct() []_domain.Product {
	var result = []_domain.Product{}
	var products []product.Product
	for _, product := range products {
		result = append(result, product.ToDomain())
	}
	return result
}

func (c *Category) ToDomainProducts(products []product.Product) []_domain.Product {
	var result = []_domain.Product{}
	for _, product := range products {
		result = append(result, product.ToDomain())
	}
	return result
}
