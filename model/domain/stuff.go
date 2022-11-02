package domain

import "time"

type Task struct {
	Id             int
	Title          string
	Description    string
	Image          string
	Parent_Task_Id int
	Poin           int
	Created_At     time.Time
	Updated_at     time.Time
}
