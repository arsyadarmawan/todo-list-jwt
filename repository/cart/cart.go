package cart

import (
	"context"
	"task/helper"
	"task/model/domain"

	"gorm.io/gorm"
)

type CartRepo interface {
	Create(ctx context.Context, cart domain.Cart) domain.Cart
	Delete(ctx context.Context, cart domain.Cart)
	FindByAll(ctx context.Context, id int) []domain.Cart
}

type CartRepoImpl struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) CartRepo {
	return &CartRepoImpl{
		db: database,
	}
}

func (repo *CartRepoImpl) Create(ctx context.Context, cart domain.Cart) domain.Cart {
	result := repo.db.Create(&cart)
	cart.Id = int(result.RowsAffected)
	return cart

}

func (repo *CartRepoImpl) Delete(ctx context.Context, product domain.Cart) {
	data := Cart{}

	query := repo.db.Delete(&data, data.Id)
	helper.PanicHandlerGORM(*query)
}

func (repo *CartRepoImpl) FindByAll(ctx context.Context, id int) []domain.Cart {
	panic("Hellow")
}
