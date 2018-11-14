package middleware

import (
	"groot/models"
)

func IsLogin(ctx *Context) {
	user := ctx.Session().Get("user")
	if user == nil {
		ctx.Go(409, "请登录")
		return
	}

	ctx.Next()
}

func IsAdmin(ctx *Context) {
	user := ctx.Session().Get("user")
	if user == nil {
		ctx.Go(409, "请登录")
		return
	}

	admin, _ := user.(*models.User)

	if !admin.IsAdmin {
		ctx.Go(406, "此操作需要管理员权限")
		return
	}

	ctx.Next()
}
