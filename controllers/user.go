package controllers

import (
	// "math/rand"
	// "time"
	// "fmt"
	"groot/models"
	"groot/middleware"
)

func SignIn(ctx *middleware.Context) {
	user := new(models.User)
	user.ID = 1
	err := user.Find()
	if err != nil {
		ctx.Go(500, "登录失败")
		return
	}

	ctx.Session().Set("user", user)
	ctx.Go(user)
}
