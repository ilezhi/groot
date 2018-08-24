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
	topics, err := TopicService.GetList(true)
	if err != nil {
		ctx.JSON(iris.Map{"code": 1, "msg": "获取帖子列表失败"})
		return
	}
	
	ctx.JSON(topics)
}

/**
 * 根据id获取话题
 */
func ByID(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	
	if err != nil {
		ctx.JSON(iris.Map{"code": 1, "msg": "参数有误"})
		return
	}

	topic, err := TopicService.GetByID(uint(id))

	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(iris.Map{"code": 1, "msg": "查询话题失败"})
		return
	}

	ctx.JSON(topic)
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
	err = TopicService.Create(&topic)

	if err != nil {
		ctx.JSON(iris.Map{"code": 1, "msg": "新增失败"})
		return
	}

	ctx.JSON(topic)
	fmt.Println("continue", topic)
}
