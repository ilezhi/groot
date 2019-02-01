package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"groot/config"
)

func init() {
	fmt.Println("执行conn.go init")
}

var DB *gorm.DB

func Connect() (db *gorm.DB, err error) {
	str := config.Values().Get("db")
	fmt.Println("链接信息", str)
	db, err = gorm.Open("mysql", str)
	if err != nil {
		fmt.Println("fail to connect database", err)
	}

	DB = db
	return
}
