package repository

import (
	_user "task/model/domain"
)

type User struct {
	Id       uint   `gorm:"primaryKey autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (user *User) ToDomain() _user.User {
	return _user.User{
		Username: user.Username,
		Name:     user.Name,
		Id:       user.Id,
		Email:    user.Email,
	}
}
