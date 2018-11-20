package controllers

import (
	"groot/models"
	"groot/middleware"
	"groot/tools"
	"groot/config"
)

type LoginInfo struct {
	Email	string			`json:"email"`
	Password	string	`json:"password"`
}

func SignIn(ctx *middleware.Context) {
	params := new(LoginInfo)
	ctx.ReadJSON(params)

	if params.Email == "" || params.Password == "" {
		ctx.Error(422, "账号或密码不能为空")
		return
	}

	// 根据邮箱获取用户
	user := new(models.User)
	user.Email = params.Email
	
	err := user.FindByEmail()
	if err != nil {
		ctx.Error(404, "用户不存在")
		return
	}

	// 验证密码是否正确
	hash := tools.EncryptPwd(params.Password)
	if user.Password != hash {
		ctx.Error(422, "用户名或密码不正确")
		return
	}

	admin := config.Values().Get("admin")
	if admin == user.Email {
		user.IsAdmin = true
	}

	ctx.Session().Set("user", user)
	data := tools.StructToMap(*user)
	delete(data, "password")
	delete(data, "token")
	delete(data, "secretKey")

	ctx.Go(data)
}

func LoginUser(ctx *middleware.Context) {
	user := ctx.Session().Get("user").(*models.User)
	data := tools.StructToMap(*user)
	delete(data, "password")
	delete(data, "token")
	delete(data, "secretKey")
	ctx.Go(data)
}

// 获取用户category, tag
func UserInfo(ctx *middleware.Context) {
	user := ctx.Session().Get("user").(*models.User)
	category := new(models.Category)
	tag := new(models.TopicTag)

	category.UserID = user.ID
	// 获取
	categories, err := category.GroupBy()
	if err != nil {
		ctx.Error(500, "获取分类失败")
		return
	}

	// 获取tags
	tags, err := tag.GroupBy(user.ID)
	if err != nil {
		ctx.Error(500, "获取标签分类失败")
		return
	}

	data := make(map[string]interface{})
	data["categories"] = categories
	data["tags"] = tags
	ctx.Go(data)
}
