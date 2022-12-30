package product

import (
	"context"
	"task/helper"
	"task/model/domain"

	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(ctx context.Context, product domain.Product) domain.Product
	Update(ctx context.Context, product domain.Product) domain.Product
	Delete(ctx context.Context, product domain.Product)
	FindById(ctx context.Context, id int) (domain.Product, error)
	FindAll(ctx context.Context) []domain.Product
	FindByName(ctx context.Context, name string) []domain.Product
}

type ProductRepoImpl struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) ProductRepo {
	return &ProductRepoImpl{
		db: database,
	}
}

func (repo ProductRepoImpl) FindByName(ctx context.Context, name string) []domain.Product {
	var products []Product

	result := repo.db.Where("name LIKE ?", "%"+name+"%").Find(&products)
	if name == "" {
		result = repo.db.Find(&products)
	}
	helper.PanicHandlerGORM(*result)
	var resultProducts []domain.Product
	for _, result := range products {
		resultProducts = append(resultProducts, domain.Product{
			Id:          int(result.Id),
			Name:        result.Name,
			Description: result.Description,
			Price:       result.Price,
			Stock:       int(result.Price),
		})
	}

	return resultProducts
}

func (repo *ProductRepoImpl) Create(ctx context.Context, product domain.Product) domain.Product {
	result := repo.db.Create(&product)
	product.Id = int(result.RowsAffected)
	return product
}

func (repo *ProductRepoImpl) Update(ctx context.Context, product domain.Product) domain.Product {
	stuff := Product{
		Name:        product.Name,
		Stock:       product.Stock,
		Category_id: int(product.Category_id),
		Price:       product.Price,
		Updated_at:  product.Updated_at,
		Description: product.Description,
	}

	query := repo.db.Model(&Product{}).Where("id = ?", stuff.Id).Updates(stuff).Find(&Product{})
	helper.PanicHandlerGORM(*query)
	return product
}

func (repo *ProductRepoImpl) Delete(ctx context.Context, product domain.Product) {
	stuff := Product{}

	query := repo.db.Delete(&stuff, product.Id)
	helper.PanicHandlerGORM(*query)
}

func (repo *ProductRepoImpl) FindById(ctx context.Context, id int) (domain.Product, error) {
	var stuff Product
	result := repo.db.First(&stuff, id)

	helper.PanicHandlerGORM(*result)
	return stuff.ToDomain(), nil
}

func (repo *ProductRepoImpl) FindAll(ctx context.Context) []domain.Product {
	panic("not implemented") // TODO: Implement
}
