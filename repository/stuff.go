package repository

import (
	"context"
	"task/helper"
	"task/model/domain"

	"gorm.io/gorm"
)

type StuffRepo interface {
	CreateStuff(ctx context.Context, stuff domain.Task) domain.Task
	Update(ctx context.Context, stuff domain.Task) domain.Task
	Delete(ctx context.Context, stuff domain.Task)
	FindById(ctx context.Context, id int) (domain.Task, error)
	FindTask(ctx context.Context, userId int) []domain.Task
}

type StuffRepoImpl struct {
	db *gorm.DB
}

func NewStuffRepository(database *gorm.DB) StuffRepo {
	return &StuffRepoImpl{
		db: database,
	}
}

func (repository *StuffRepoImpl) CreateStuff(ctx context.Context, stuff domain.Task) domain.Task {
	task := Task{
		Created_At: stuff.Created_At,
		Updated_at: stuff.Updated_at,
		Product_id: stuff.Product_id,
		User_id:    stuff.User_id,
		Total:      stuff.Total,
	}

	result := repository.db.Create(&task)
	helper.PanicHandling(result.Error)
	stuff.Id = int(result.RowsAffected)
	return stuff
}

func (repository *StuffRepoImpl) Update(ctx context.Context, stuff domain.Task) domain.Task {
	task := Task{
		Total: stuff.Total,
	}
	query := repository.db.Model(&Task{}).Where("product_id = ?", stuff.Id).Updates(task).Find(&Task{})
	helper.PanicHandlerGORM(*query)
	return stuff
}

func (repository *StuffRepoImpl) Delete(ctx context.Context, stuff domain.Task) {
	task := Task{}

	query := repository.db.Delete(&task, stuff.Id)
	helper.PanicHandlerGORM(*query)
}

func (repository *StuffRepoImpl) FindById(ctx context.Context, id int) (domain.Task, error) {
	var task Task
	result := repository.db.First(&task, id)

	helper.PanicHandlerGORM(*result)
	return task.ToDomain(), nil
}

func (repository *StuffRepoImpl) FindTask(ctx context.Context, userId int) []domain.Task {
	var tasks []Task
	result := repository.db.Where("user_id = ? ", userId).Find(&tasks)

	helper.PanicHandlerGORM(*result)
	var resultTask []domain.Task
	for _, result := range tasks {
		resultTask = append(resultTask, domain.Task{
			Created_At: result.Created_At,
			Updated_at: result.Updated_at,
			Product_id: result.Product_id,
			Total:      result.Total,
			Id:         int(result.Id),
			User_id:    result.User_id,
		})
	}

	return resultTask
}
