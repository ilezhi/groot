package main

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/kataras/iris"
	// "github.com/kataras/iris/middleware/logger"

	. "groot/model"
	// "groot/route"
)

func init() {
	db, err := gorm.Open("mysql", "root:Mysql@2018@/groot?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		os.Exit(2)
	}

	err = db.CreateTable(&User{}, &Topic{}, &Tag{}, &TopicTag{}, &Comment{}, &Reply{}, &Project{}, 
		&ProjectMember{}, &Good{}, &Favor{}, &Department{}, &Category{}).Error

	if err != nil {
		fmt.Println("创建表格失败", err)
		os.Exit(2)
	}
}

func main() {
	// app := iris.New()
	
	// app.Use(logger.New())
	
	// Router.Register(app)

	// app.Run(
	// 	iris.Addr("localhost:9000"),
	// 	iris.WithoutVersionChecker,
	// 	iris.WithoutServerError(iris.ErrServerClosed),
	// 	iris.WithOptimizations,
	// )
}
