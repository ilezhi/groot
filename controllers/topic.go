package topic

import (
	"fmt"
	"github.com/kataras/iris"

	// . "groot/db"
	. "groot/model"
	. "groot/services"
)

/**
 * 登录
 */
func List(ctx iris.Context) {
	ctx.JSON(iris.Map{"name": "topic list"})
}

func ByID(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	ctx.JSON(iris.Map{"topicID": id})
}

func Create(ctx iris.Context) {
	var topic Topic
	err := ctx.ReadJSON(&topic)

	if err != nil {
		ctx.JSON(iris.Map{"code": 1, "msg": "无法序列化成json格式"})
		return
	}

	topic.AuthorID= 10000
	err = topic.Validate()

	if err != nil {
		ctx.JSON(iris.Map{"code": 1, "msg": "参数验证失败"})
		return
	}

	// 添加话题和tag
	err = TS.Create(&topic)

	if err != nil {
		ctx.JSON(iris.Map{"code": 1, "msg": "新增失败"})
		return
	}

	ctx.JSON(topic)
	fmt.Println("continue", topic)
}
