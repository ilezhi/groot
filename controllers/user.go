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

type SignupInfo struct {
	LoginInfo
	ConfirmPassword string `json:"confirmPassword"`
	Nickname   string      `json:"nickname"`
	DeptID		 uint				 `json:"deptID"`
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

func SignUp(ctx *middleware.Context) {
	params := new(SignupInfo)
	ctx.ReadJSON(params)

	if params.Email == "" || params.Password == "" {
		ctx.Error(422, "账号或密码不能为空")
		return
	}

	if params.Password != params.ConfirmPassword {
		ctx.Error(422, "两次密码输入不一致")
		return
	}

	user := new(models.User)
	user.Email = params.Email
	err := user.FindByEmail()
	if err == nil {
		ctx.Error(422, "邮箱已存在")
		return
	}

	user.Password = tools.EncryptPwd(params.Password)
	user.Nickname = params.Nickname
	user.Name = params.Nickname
	user.Avatar = tools.GetAvatar(params.Email)
	user.DeptID = params.DeptID
	err = user.Save()
	if err != nil {
		ctx.Error(500, "注册失败")
		return
	}

	
	ctx.Go("注册成功")
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

func IsExistByEmail(ctx *middleware.Context) {
	email := ctx.URLParam("email")
	if email == "" {
		ctx.Error(400)
		return
	}
	
	user := new(models.User)
	user.Email = email
	err := user.FindByEmail()
	if err == nil {
		ctx.Go(true)
		return
	}

	ctx.Go(false)
}

func Departments(ctx *middleware.Context) {
	dept := new(models.Department)
	depts, err := dept.List()
	if err != nil {
		ctx.Error(500)
		return
	}

	ctx.Go(depts)
}
