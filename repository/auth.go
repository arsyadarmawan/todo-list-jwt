package repository

import (
	"fmt"
	"task/helper"
	"task/model/domain"

	"gorm.io/gorm"
)

type AuthRepo interface {
	CheckUsername(username string) domain.User
	Register(user *domain.User) error
	Login(username string) (string, error)
	CheckId(id int) bool
	Delete(id int) error
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &AuthRepository{db: db}
}
func (r AuthRepository) CheckUsername(username string) domain.User {
	fmt.Println(username)
	var user User
	result := r.db.Where("username = ?", username).First(&user)

	helper.PanicHandlerGORM(*result)
	return user.ToDomain()
}

func (r AuthRepository) Register(user *domain.User) error {
	if err := r.db.Table("users").Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r AuthRepository) Login(username string) (string, error) {
	var user domain.User
	if err := r.db.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	return user.Password, nil
}

func (r AuthRepository) CheckId(id int) bool {
	var count int64
	if err := r.db.Table("users").Where("id = ?", id).Count(&count).Error; err != nil {
		return false
	}

	if count < 1 {
		return false
	}

	return true
}

func (r AuthRepository) Delete(id int) error {
	if err := r.db.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
