package web

import "time"

type TaskResponse struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Image          string    `json:"image"`
	Description    string    `json:"description"`
	Parent_Task_Id int       `json:"parent_task_id"`
	Poin           int       `json:"poin"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}
