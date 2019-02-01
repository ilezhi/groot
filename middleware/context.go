package middleware

import (
	"time"
	"sync"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"

	"groot/models"
	"groot/config"
)

var cookieNameForSessionID = "mycookiesessionnameid"

type Owner struct {
	createSess		*sessions.Sessions
}

var owner = &Owner{
	createSess: sessions.New(sessions.Config{
		Cookie: cookieNameForSessionID,
		Expires: 10 * time.Hour,
		AllowReclaim: true,
	}),
}

type Context struct {
	iris.Context
	sess					*sessions.Session
	client				*Client
	Config				*config.Config
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original
	ctx.sess = nil
	ctx.client = nil
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

func (ctx *Context) Session() *sessions.Session {
	if ctx.sess == nil {
		ctx.sess = owner.createSess.Start(ctx.Context)
	}
	return ctx.sess
}

func (ctx *Context) Client() *Client {
	if ctx.client == nil {
		user := ctx.Session().Get("user").(*models.User)
		ctx.client = hub.clients[user.ID]
	}

	return ctx.client
}

func (ctx *Context) Go(p interface{}) {
	code := 0
	message := "请求成功"
	data := p

	ctx.Values().Set("code", code)
	ctx.Values().Set("message", message)
	ctx.Values().Set("data", data)
}
