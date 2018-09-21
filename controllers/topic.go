package controllers

import (
	"fmt"
	"time"
	"groot/models"
	. "groot/services"
	"groot/tools"
)

/**
 * 获取topic list
 */
func Topics(ctx *tools.Context) {
	lastID, _ := ctx.URLParamInt("lastID")
	size, _ := ctx.URLParamInt("size")

	topics, err := TopicService.Find(uint(lastID), size)
	if err != nil {
		ctx.Go(1, "获取帖子列表失败")
		return
	}

	ctx.Go(topics)
}

/**
 * 根据id获取话题
 */
func Topic(ctx *tools.Context) {
	id, err := ctx.Params().GetInt("id")
	
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	topic, err := TopicService.FindByID(uint(id))

	if err != nil {
		ctx.Go(500, "查询话题失败")
		return
	}

	topic.View++
	topic.UpdateView()

	ctx.Go(topic)
}

/**
 * 新增
 */
func CreateTopic(ctx *tools.Context) {
	var params models.TopicParams
	var topic models.Topic
	err := ctx.ReadJSON(&params)

	if err != nil {
		ctx.Go(1, "参数有误")
		return
	}
	fmt.Println("params", params)
	topic.Title = params.Title
	topic.Content = params.Content
	topic.Shared = params.Shared
	topic.AuthorID= 10000
	topic.UpdatedAt = time.Now().Unix()
	err = topic.Validate()

	if err != nil {
		ctx.Go(1, "参数验证失败")
		return
	}

	// 添加话题和tag
	err = TopicService.Create(&topic, &params.Tags)

	if err != nil {
		ctx.Go(1, "新增失败")
		return
	}

	topic.Tags, err = TagService.FindByIDs(params.Tags)

	if err != nil {
		ctx.Go(500, "获取tag失败")
		return
	}

	ctx.Go(topic)
}

/**
 * 更新
 */
func UpdateTopic(ctx *tools.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	// 根据id获取topic
	var params models.TopicParams
	err = ctx.ReadJSON(&params)
	topic, err := TopicService.FindAndUpdate(uint(id), params.Content, &params.Tags)

	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}

	ctx.Go(topic)
}

func TrashTopic(ctx *tools.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	err = TopicService.DeleteByID(uint(id))
	if err != nil {
		ctx.Go(500, "删除失败")
		return
	}

	ctx.Go(id)
}

/**
 * 精华
 */
func AwesomeTopic(ctx *tools.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	params := map[string]bool{"awesome": false}
	err = ctx.ReadJSON(&params)

	if err != nil {
		ctx.Go(406, "body 参数有误")
		return
	}

	topic, err := TopicService.FindAndUpdateColumns(uint(id), params)

	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}

	ctx.Go(topic)
}

/**
 * 分享
 */
func ShareTopic(ctx *tools.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	params := map[string]bool{"shared": false}
	err = ctx.ReadJSON(&params)
	if err != nil {
		ctx.Go(406, "body 参数有误")
		return
	}

	topic, err := TopicService.FindAndUpdateColumns(uint(id), params)
	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}
	
	ctx.Go(topic)
}

/**
 * 置顶
 */
func TopTopic(ctx *tools.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	params := map[string]bool{"top": false}
	err = ctx.ReadJSON(&params)
	if err != nil {
		ctx.Go(406, "body 参数有误")
		return
	}

	topic, err := TopicService.FindAndUpdateColumns(uint(id), params)
	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}

	ctx.Go(topic)
}

/**
 * 收藏
 */
func FavorTopic(ctx *tools.Context) {

}