package services

import (
	sql "groot/db"
	. "groot/models"
)

type IUser interface {
	FindByID(id uint) (*User, error)
	Create(user *User) error
	CreateCategory(category *Category) error
}

type userService struct {}

var UserService = userService{}


func (us * userService) FindByID(id uint) (*User, error) {
	var user User
	err := sql.DB.First(&user, id).Error

	return &user, err
}

func (us *userService) Create(user *User) error {
	return sql.DB.Create(user).Error 
}

func (us *userService) CreateCategory(category *Category) error {
	return sql.DB.Create(category).Error
}
