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
	db, err = gorm.Open("mysql", config.Values().Get("db"))
	if err != nil {
		fmt.Println("fail to connect database", err)
	}

	DB = db
	return
}
