package topic

import (
	"fmt"
	"github.com/kataras/iris"

	. "groot/model"
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
	if err := ctx.ReadJSON(&topic); err != nil {
		ctx.JSON(iris.Map{"code": 1, "msg": "参数有误"})
		return
	}

	ctx.JSON(topic)
	fmt.Println("continue", topic)
}
