package domain

import "time"

type Cart struct {
	Id         int
	User_id    int
	Product_id int
	Created_At time.Time
	Updated_at time.Time
}
