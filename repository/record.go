package repository

import (
	_task "task/model/domain"
	"time"
)

type Task struct {
	Id             uint `gorm:"primaryKey autoIncrement"`
	Title          string
	Description    string
	Image          string
	Parent_Task_Id int
	Poin           int
	Created_At     time.Time
	Updated_at     time.Time
}

func (task *Task) ToDomain() _task.Task {
	return _task.Task{
		Id:             int(task.Id),
		Title:          task.Title,
		Description:    task.Description,
		Image:          task.Image,
		Created_At:     task.Created_At,
		Updated_at:     task.Updated_at,
		Poin:           task.Poin,
		Parent_Task_Id: int(task.Parent_Task_Id),
	}
}
