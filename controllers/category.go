package controllers

import (
	// "math/rand"
	// "time"
	// "fmt"
	"groot/models"
	"groot/middleware"
)

func CreateCategory(ctx *middleware.Context) {
	category := new(models.Category)
	ctx.ReadJSON(category)
	user := ctx.Session().Get("user").(*models.User)
	category.UserID = user.ID

	isExist := category.IsExist()
	if isExist {
		ctx.Error(403, "资源已存在")
		return
	}

	err := category.Save()
	if err != nil {
		ctx.Error(500, "新增分类失败")
		return
	}

	ctx.Go(category)
}
