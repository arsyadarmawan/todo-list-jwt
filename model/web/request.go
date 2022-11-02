package web

type TaskCreateRequest struct {
	Title          string `validate:"required,min=1,max=100" json:"title"`
	Description    string `validate:"required,min=1,max=1000" json:"description"`
	Image          string `validate:"min=1,max=150" json:"image" contains:".jpg"`
	Poin           int    `validate:"required,min=1,max=5" json:"poin"`
	Parent_Task_Id int    `validate:"omitempty" json:"parent_task_id"`
}

type TaskUpdateRequest struct {
	Id             int    `validate:"required"`
	Title          string `validate:"required,min=1,max=100" json:"title"`
	Description    string `validate:"required,min=1,max=1000" json:"description"`
	Image          string `validate:"min=1,max=150" json:"image" contains:".jpg"`
	Poin           int    `validate:"required" json:"poin"`
	Parent_Task_Id int    `json:"parent_task_id"`
}
