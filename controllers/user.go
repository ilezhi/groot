package controllers

import (
	"fmt"
	"groot/models"
	. "groot/services"
	"groot/middleware"
)

/**
 * 新增用户
 */
func CreateUser(ctx *middleware.Context) {
	var user models.User

	err := ctx.ReadJSON(&user)

	if err != nil {
		ctx.Go(406, "参数错误")
		return
	}

	err = UserService.Create(&user)

	if err != nil {
		ctx.Go(500, "新增用户失败")
		return
	}

	ctx.Go(user)
}

func Login(ctx *middleware.Context) {
	id := uint(4)
	user, err := UserService.FindByID(id)
	if err != nil {
		fmt.Println("err", err)
		ctx.Go(500, "用户不存在")
		return
	}

	ctx.Session().Set("user", user)
	ctx.Go(user)
}

/**
 * 新建收藏文件夹
 */
func CreateCategory(ctx *middleware.Context) {
	var category models.Category
	
	err := ctx.ReadJSON(&category)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	category.UserID = user.ID

	err = UserService.CreateCategory(&category)
	if err != nil {
		ctx.Go(500, "新建分类失败")
		return
	}

	ctx.Go(category)
}
