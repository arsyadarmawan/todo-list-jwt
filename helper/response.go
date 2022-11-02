package helper

import (
	"task/model/domain"
	"task/model/web"
)

func StuffResponse(stuff domain.Task) web.TaskResponse {
	return web.TaskResponse{
		Id:             stuff.Id,
		Title:          stuff.Title,
		Description:    stuff.Description,
		Created_at:     stuff.Created_At,
		Updated_at:     stuff.Updated_at,
		Image:          stuff.Image,
		Parent_Task_Id: int(stuff.Parent_Task_Id),
		Poin:           int(stuff.Poin),
	}
}
