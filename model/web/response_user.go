package web

import "time"

type UserReponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
