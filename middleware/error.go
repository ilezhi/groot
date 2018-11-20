package middleware

var errors = map[int]string{
	400: "无法识别请求参数",
	401: "请登录",
	403: "需要管理员权限",
	404: "帖子不存在",
	410: "资源已删除",
	422: "参数验证失败",
	423: "资源被锁定",
	500: "数据库操作失败",
}

func (ctx *Context) Error(p ...interface{}) {
	var code int
	var message string

	n := len(p)
	code, _ = p[0].(int)

	if n == 1 {
		message = errors[code]
	}

	if n == 2 {
		message, _ = p[1].(string) 
	}

	ctx.Values().Set("code", code)
	ctx.Values().Set("message", message)
}
