package middleware

import (
	"groot/models"
)

func IsLogin(ctx *Context) {
	user := ctx.Session().Get("user")
	if user == nil {
		ctx.Error(401)
		return
	}

	ctx.Next()
}

func IsAdmin(ctx *Context) {
	user := ctx.Session().Get("user")
	if user == nil {
		ctx.Go(401)
		return
	}

	admin, _ := user.(*models.User)

	if !admin.IsAdmin {
		ctx.Go(403)
		return
	}

	ctx.Next()
}
