package domain

import (
	"time"
)

type Task struct {
	Id         int
	Product_id int
	User_id    int
	Total      int
	Created_At time.Time
	Updated_at time.Time
}
