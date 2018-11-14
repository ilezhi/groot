package middleware

import (
	"groot/config"
)

func SetConfig(ctx *Context) {
	ctx.Config = config.Values()
	ctx.Next()
}
