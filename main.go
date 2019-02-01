package main

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"groot/db"
	"groot/route"
	"groot/middleware"
	"groot/config"
	. "groot/models"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		return
	}

	if !conn.HasTable("users") {
		initDB(conn)
	}

	defer conn.Close()
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": ctx.StatusCode, "message": ctx.Values().GetString("message")})
	})

	app.Use(middleware.Handler(middleware.SetConfig))

	app.Use(middleware.Response)
	app.Get("/ws/{id:int}", middleware.Handler(middleware.WSConn))

	Router.Register(app)

	app.Run(
		iris.Addr(config.Values().Get("localhost").(string)),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func initDB(DB *gorm.DB) {
	err := DB.CreateTable(&User{}, &Topic{}, &Tag{}, &TopicTag{}, &Comment{}, &Reply{}, &Project{}, 
		&ProjectMember{}, &Like{}, &Favor{}, &Department{}, &Category{}, &Team{}).Error
	if err != nil {
		fmt.Println("fail to create tables")
		os.Exit(2)
	}

	depts := config.Values().Get("departments")
	arr, _ := depts.([]string)
	for _, val := range arr {
		dept := &Department{
			Name: val,
		}
		err = DB.Create(dept).Error
		if err != nil {
			fmt.Println("fail to create departments")
		}
	}
}
