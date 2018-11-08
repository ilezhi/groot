package middleware

import (
	"fmt"
)

func IsLogin(ctx *Context) {
	fmt.Println("path", ctx.Path())
	user := ctx.Session().Get("user")
	if user == nil {
		ctx.Go(407, "请登录")
		return
	}

	ctx.Next()
}