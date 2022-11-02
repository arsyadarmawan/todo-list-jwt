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
	FindTask(ctx context.Context) []domain.Task
	FindTaskByTaskId(ctx context.Context, id int) []domain.Task
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
		Title:       stuff.Title,
		Description: stuff.Description,
		Image:       stuff.Image,
		// Parent_Task_Id: stuff.Parent_Task_Id,
		Poin: stuff.Poin,
	}

	result := repository.db.Create(&task)
	helper.PanicHandling(result.Error)
	stuff.Id = int(result.RowsAffected)
	return stuff
}

func (repository StuffRepoImpl) Update(ctx context.Context, stuff domain.Task) domain.Task {
	task := Task{
		Title:          stuff.Title,
		Description:    stuff.Description,
		Parent_Task_Id: stuff.Parent_Task_Id,
		Image:          stuff.Image,
	}
	query := repository.db.Model(&Task{}).Where("id = ?", stuff.Id).Updates(task).Find(&Task{})
	helper.PanicHandlerGORM(*query)
	return stuff
}

func (repository StuffRepoImpl) Delete(ctx context.Context, stuff domain.Task) {
	task := Task{}

	query := repository.db.Delete(&task, stuff.Id)
	helper.PanicHandlerGORM(*query)
}

func (repository StuffRepoImpl) FindById(ctx context.Context, id int) (domain.Task, error) {
	var task Task
	result := repository.db.First(&task, id)

	helper.PanicHandlerGORM(*result)
	return task.ToDomain(), nil
}

func (repository StuffRepoImpl) FindTask(ctx context.Context) []domain.Task {
	var tasks []Task
	result := repository.db.Where("parent_task_id = ? ", 0).Order("poin desc").Find(&tasks)

	helper.PanicHandlerGORM(*result)
	var resultTask []domain.Task
	for _, result := range tasks {
		resultTask = append(resultTask, domain.Task{
			Id:             int(result.Id),
			Title:          result.Title,
			Description:    result.Description,
			Poin:           result.Poin,
			Image:          result.Image,
			Parent_Task_Id: result.Parent_Task_Id,
		})
	}

	return resultTask
}

func (repository StuffRepoImpl) FindTaskByTaskId(ctx context.Context, id int) []domain.Task {
	var tasks []Task
	result := repository.db.Where("parent_task_id = ?", id).Order("poin desc").Find(&tasks)

	helper.PanicHandlerGORM(*result)
	var resultTask []domain.Task
	for _, result := range tasks {
		resultTask = append(resultTask, domain.Task{
			Id:             int(result.Id),
			Title:          result.Title,
			Description:    result.Description,
			Image:          result.Image,
			Poin:           result.Poin,
			Parent_Task_Id: result.Parent_Task_Id,
		})
	}
	return resultTask
}
