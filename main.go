package main

import (
	// "fmt"
	// "os"
	// "github.com/jinzhu/gorm"
	// _ "github.com/go-sql-driver/mysql"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	// . "groot/models"
	// "groot/tools"

	"groot/db"
	"groot/route"
	"groot/middleware"
)

// var DB *gorm.DB

// func init() {
// 	DB, err := gorm.Open("mysql", "root:Mysql@2018@/groot?charset=utf8&parseTime=True&loc=Local")
// 	defer DB.Close()
// 	if err != nil {
// 		fmt.Println("连接数据库失败", err)
// 		os.Exit(2)
// 	}

	// if err != nil {
	// 	fmt.Println("验证失败", err)
	// } else {
	// 	fmt.Println("验证成功")
	// }
	// err = DB.CreateTable(&User{}, &Topic{}, &Tag{}, &TopicTag{}, &Comment{}, &Reply{}, &Project{}, 
	// 	&ProjectMember{}, &Like{}, &Favor{}, &Department{}, &Category{}, &Team{}).Error

	// if err != nil {
	// 	fmt.Println("创建表格失败", err)
	// 	os.Exit(2)
	// }

// 	hash := tools.EncryptPwd("123abc")
// 	user := User{
// 		Name: "skr",
// 		Nickname: "skr",
// 		Email: "skr@123.com",
// 		Password: hash,
// 		IsAdmin: false,
// 		IsVerify: true,
// 	}

// 	err = DB.Create(&user).Error

// 	if err != nil {
// 		fmt.Println("新增用户错误", err)
// 	}
// 	fmt.Println(user)
// }

func main() {
	conn, _ := db.Connect()
	defer conn.Close()
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": ctx.StatusCode, "message": ctx.Values().GetString("message")})
	})

	app.Use(middleware.Response)
	app.Get("/ws/{id:int}", middleware.Handler(middleware.WSConn))

	Router.Register(app)

	app.Run(
		iris.Addr("172.18.2.231:9000"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
