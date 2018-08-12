package main

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	. "groot/model"
)

var engine *xorm.Engine

func main() {
	engine, err := xorm.NewEngine("mysql", "root:mmbeibei@/groot?charset=utf8")
	defer engine.Close()
	if err != nil {
		fmt.Println("连接失败")
		os.Exit(1)
	}

	exist, _ := engine.IsTableExist(new(User))

	if !exist {
		err = engine.CreateTables(new(User), new(Role))

		if err != nil {
			fmt.Println("fail to create table user")
			os.Exit(2)
		}
	}

	// role := new(Role)
	// role.Name = "普通用户"
	// engine.Insert(role)

	user := new(User)
	user.Name = "Stark"
	user.Birthday = "1990-11-09"
	affect, er := engine.Insert(user)

	if er != nil {
		fmt.Println("fail to insert user")
	} else {
		fmt.Println("success", affect)
	}
}