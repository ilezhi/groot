package tools

import (
	"sync"
	"github.com/kataras/iris"
)

type Context struct {
	iris.Context
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original // set the context to the original one in order to have access to iris's implementation.
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

func Handler(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}

func (ctx *Context) Go(a ...interface{}) {
	n := len(a)
	var code int
	var message string
	var data interface{}

	if n == 1 {
		code = 0
		message = "请求成功"
		data = a[0]
	}
	
	if n == 2 {
		code, _ = a[0].(int)
		message, _ = a[1].(string)
	}

	ctx.Values().Set("code", code)
	ctx.Values().Set("message", message)
	ctx.Values().Set("data", data)
}
