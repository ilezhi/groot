package services

import (
	mysql "groot/db"
	. "groot/models"
)

type IUser interface {
	Create(user *User) error 
}

type userService struct {}

var UserService = userService{}

func (us *userService) Create(user *User) error {
	return mysql.DB.Create(user).Error 
}
