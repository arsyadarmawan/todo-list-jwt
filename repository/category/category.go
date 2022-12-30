package category

import (
	"context"
	"task/helper"
	"task/model/domain"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(ctx context.Context, category domain.Category) domain.Category
	Delete(ctx context.Context, category domain.Category)
	FindById(ctx context.Context, id int) (domain.Category, error)
	FindByAll(ctx context.Context, name string) []domain.Category
}

func NewCategoryRepository(database *gorm.DB) CategoryRepo {
	return &CategoryRepoImpl{
		db: database,
	}
}

type CategoryRepoImpl struct {
	db *gorm.DB
}

func (repo *CategoryRepoImpl) Create(ctx context.Context, category domain.Category) domain.Category {
	result := repo.db.Create(&category)
	category.Id = int(result.RowsAffected)
	return category
}

func (repo *CategoryRepoImpl) Delete(ctx context.Context, category domain.Category) {
	data := Category{}

	query := repo.db.Delete(&data, category.Id)
	helper.PanicHandlerGORM(*query)
}

func (repo *CategoryRepoImpl) FindById(ctx context.Context, id int) (domain.Category, error) {
	var category Category
	err := repo.db.Model(&Category{}).Preload("Products").Find(&category).Error
	if err != nil {
		panic("Error")
	}
	return category.ToDomain(), nil
}

func (repo *CategoryRepoImpl) FindByAll(ctx context.Context, name string) []domain.Category {
	var categories []Category
	err := repo.db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic("Error")
	}

	var resultTask []domain.Category
	for _, result := range categories {
		productList := result.ToDomainProducts(result.Products)
		resultTask = append(resultTask, domain.Category{
			Id:         int(result.Id),
			Name:       result.Name,
			Products:   productList,
			Created_At: result.Created_At,
			Updated_at: result.Updated_at,
		})
	}

	return resultTask
}
