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
		ctx.Go(406, "分类已存在")
		return
	}

	err := category.Save()
	if err != nil {
		ctx.Go(500, "新增失败")
		return
	}

	ctx.Go(category)
}
