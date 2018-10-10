package controllers

import (
	"groot/models"
	"groot/middleware"
	"groot/tools"
)

type LoginInfo struct {
	Email	string			`json:"email"`
	Password	string	`json:"password"`
}

func SignIn(ctx *middleware.Context) {
	params := new(LoginInfo)
	ctx.ReadJSON(params)

	if params.Email == "" || params.Password == "" {
		ctx.Go("406", "账号或密码不能为空")
		return
	}

	// 根据邮箱获取用户
	user := new(models.User)
	user.Email = params.Email
	
	err := user.FindByEmail()
	if err != nil {
		ctx.Go(500, "用户不存在")
		return
	}

	// 验证密码是否正确
	hash := tools.EncryptPwd(params.Password)
	if user.Password != hash {
		ctx.Go(406, "用户名或密码不正确")
		return
	}

	ctx.Session().Set("user", user)
	data := tools.StructToMap(*user)
	delete(data, "password")
	delete(data, "token")
	delete(data, "secretKey")
	ctx.Go(data)
}

// 获取用户category, tag
func UserInfo(ctx *middleware.Context) {
	user := ctx.Session().Get("user").(*models.User)
	favor := new(models.Favor)
	tag := new(models.TopicTag)

	favor.UserID = user.ID
	// 获取
	categories, err := favor.GroupByCategory()
	if err != nil {
		ctx.Go(500, "获取分类失败")
		return
	}

	// 获取tags
	tags, err := tag.GroupByTag(user.ID)
	if err != nil {
		ctx.Go(500, "获取标签分类失败")
		return
	}

	data := make(map[string]interface{})
	data["categories"] = categories
	data["tags"] = tags
	ctx.Go(data)
}
