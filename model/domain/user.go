package domain

import (
	"time"
)

type GormCustom struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	GormCustom
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Id       uint   `json:"id" gorm:"primary_key"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
