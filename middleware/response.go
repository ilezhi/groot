package middleware

import (
	"github.com/kataras/iris"
)

func Response(ctx iris.Context)  {
	ctx.Next()

	code, _ := ctx.Values().GetInt("code")
	message := ctx.Values().GetString("message")
	data := ctx.Values().Get("data")
	ctx.JSON(iris.Map{"code": code, "message": message, "data": data})
}
