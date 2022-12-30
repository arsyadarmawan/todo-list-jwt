package repository

import (
	_task "task/model/domain"
	"time"
)

type Task struct {
	Id         uint `gorm:"primaryKey autoIncrement"`
	Product_id int
	User_id    int
	Total      int
	Created_At time.Time
	Updated_at time.Time
}

func (task *Task) ToDomain() _task.Task {
	return _task.Task{
		Id:         int(task.Id),
		Total:      task.Total,
		Product_id: int(task.Product_id),
		Created_At: task.Created_At,
		Updated_at: task.Updated_at,
	}
}
