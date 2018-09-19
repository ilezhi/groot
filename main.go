package main

import (
	"fmt"
	// "os"
	// "github.com/jinzhu/gorm"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	// . "groot/models"
	"groot/db"
	"groot/route"
)

// var DB *gorm.DB

// func init() {
// 	DB, err := gorm.Open("mysql", "root:Mysql@2018@/groot?charset=utf8&parseTime=True&loc=Local")
// 	defer DB.Close()
// 	if err != nil {
// 		fmt.Println("连接数据库失败", err)
// 		os.Exit(2)
// 	}

// 	user := User{
// 		BaseModel: BaseModel{ID: 10000}
// 		Name: "董明",
// 		Nickname: "",
// 		Email: "dongmingming@renrenche.com",
// 		IsAdmin: true,
// 		IsVerify: true,
// 	}

// 	err = DB.Create(&user).Error

// 	if err != nil {
// 		fmt.Println("新增用户错误", err)
// 	}
// 	fmt.Println(user)

// 	err := user.Validate()
// 	if err != nil {
// 		fmt.Println("验证失败", err)
// 	} else {
// 		fmt.Println("验证成功")
// 	}
	// err = DB.CreateTable(&User{}, &Topic{}, &Tag{}, &TopicTag{}, &Comment{}, &Reply{}, &Project{}, 
	// 	&ProjectMember{}, &Good{}, &Favor{}, &Department{}, &Category{}, &Team{}).Error

	// if err != nil {
	// 	fmt.Println("创建表格失败", err)
	// 	os.Exit(2)
	// }
// }

func main() {
	conn, _ := db.Connect()
	defer conn.Close()
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())
	
	app.Use(func(ctx iris.Context) {
		ctx.Next()

		fmt.Println("status", ctx.GetStatusCode())
		code, _ := ctx.Values().GetInt("code")
		message := ctx.Values().GetString("message")
		data := ctx.Values().Get("data")
		ctx.JSON(iris.Map{"code": code, "message": message, "data": data})
	})

	Router.Register(app)

	app.DoneGlobal(after)

	app.Run(
		iris.Addr("localhost:9000"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func after(ctx iris.Context) {
	fmt.Println("after", ctx.Values().Get("data"))
}
