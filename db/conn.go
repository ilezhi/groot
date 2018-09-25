package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("执行conn.go init")
}

var DB *gorm.DB

func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:Mysql@2018@/groot?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("fail to connect database", err)
	}
	DB = db
	return
}
