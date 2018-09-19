package controllers

import (
	"groot/models"
	. "groot/services"
	"groot/tools"
)

/**
 * 获取topic list
 */
func TopicList(ctx *tools.Context) {
	topics, err := TopicService.GetList(true)
	if err != nil {
		ctx.Go(1, "获取帖子列表失败")
		return
	}

	ctx.Go(topics)
}

/**
 * 根据id获取话题
 */
func TopicByID(ctx *tools.Context) {
	id, err := ctx.Params().GetInt("id")
	
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	topic, err := TopicService.GetByID(uint(id))

	if err != nil {
		ctx.Go(500, "查询话题失败")
		return
	}

	ctx.Go(topic)
}

/**
 * 新增topic 
 */
func CreateTopic(ctx *tools.Context) {
	var topic models.Topic
	err := ctx.ReadJSON(&topic)

	if err != nil {
		ctx.Go(1, "参数有误")
		return
	}

	topic.AuthorID= 10000
	err = topic.Validate()

	if err != nil {
		ctx.Go(1, "参数验证失败")
		return
	}

	// 添加话题和tag
	err = TopicService.Create(&topic)

	if err != nil {
		ctx.Go(1, "新增失败")
		return
	}

	ctx.Go(topic)
}
