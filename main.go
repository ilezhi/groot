package main

import (
	"fmt"
	"os"
	// "encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	. "groot/model"
)

func main() {
	db, err := gorm.Open("mysql", "root:Mysql@2018@/groot?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println("连接数据库失败", err)
		os.Exit(2)
	}

	db.CreateTable(&User{}, &Topic{}, &Tag{}, &TopicTag{}, &Comment{}, &Good{}, &Reply{})

	// 创建用户
	// user := User{Name: "Captain", Email: "captain@gmail.com", IsAdmin: true}
	// db.Create(&user)

	// 创建tag
	// tag := Tag{Name: "javascript", Description: "编程语言", AuthorID: 2}
	// db.Create(&tag)

	// 创建topic
	// topic := Topic{Title: "如何成为一个超级英雄", Content: "<strong>洗洗睡吧</strong>", AuthorID: 1}
	// db.Create(&topic)
	// tag := TopicTag{TopicID: 2, TagID: 2}
	// db.Create(&tag)

	// 添加评论
	// comment := Comment{Content: "确实不错", AuthorID: 2, TopicID: 1}
	// db.Create(&comment)

	// 添加回复
	// reply := Reply{Content: "haha", CommentID: 1, AuthorID: 1, ReceiverID: 2}
	// db.Create(&reply)

	// 点赞
	// good := Good{UserID: 2, TargetID: 1, Type: "topic"}
	// db.Create(&good)

	// var comments []*Comment
	// rows, err := db.Raw("select c.*, u.name from comments c inner join users u on c.author_id = u.id").Rows()
	// if err != nil {
	// 	fmt.Println("查询出错", err)
	// 	os.Exit(2)
	// }

	// defer rows.Close()
	// for rows.Next() {
	// 	var comment Comment
	// 	db.ScanRows(rows, &comment)
	// 	fmt.Println(comment)
	// 	comments = append(comments, &comment)
	// }
	// db.Find(&comments)
	// if data, err := json.Marshal(comments); err == nil {
	// 	fmt.Printf("%s\n", data)
	// }
}
