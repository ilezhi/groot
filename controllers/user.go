package controllers

import (
	"groot/models"
	. "groot/services"
	"groot/tools"
)

/**
 * 新增用户
 */
func CreateUser(ctx *tools.Context) {
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
