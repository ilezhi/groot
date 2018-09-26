package controllers

import (
	"math/rand"
	"time"
	"fmt"
	"groot/models"
	. "groot/services"
	"groot/middleware"
)

/**
 * 获取topic list
 */
func Topics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")

	num := rand.Intn(10)
	fmt.Println("int", num)
	time.Sleep(time.Duration(num) * time.Second)

	query := map[string]interface{}{"issue": 1}
	topics, err := TopicService.FindByQuery(lastID, query)
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

/**
 * 精华
 */
func AwesomeTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")

	query := map[string]interface{}{"awesome": 1}
	topics, err := TopicService.FindByQuery(lastID, query)
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func MyTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")

	user := ctx.Session().Get("user").(*models.User)

	query := map[string]interface{}{"author_id": user.ID}
	topics, err := TopicService.FindByQuery(lastID, query)
	if err != nil {
		fmt.Println("err", err)
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func QTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")

	user := ctx.Session().Get("user").(*models.User)
	topics, err := TopicService.FindQuestion(user.ID, lastID)
	if err != nil {
		fmt.Println("err", err)
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func ATopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")

	user := ctx.Session().Get("user").(*models.User)
	topics, err := TopicService.FindAnswer(user.ID, lastID)
	if err != nil {
		fmt.Println("err", err)
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

/**
 * 根据id获取话题
 */
func Topic(ctx *middleware.Context) {
	id, err := ctx.Params().GetInt("id")
	
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	topic, err := TopicService.FindByID(uint(id))
	if err != nil {
		ctx.Go(500, "话题不存在")
		return
	}

	topic.View++
	topic.UpdateView()

	ctx.Go(topic)
}

/**
 * 新增
 */
func CreateTopic(ctx *middleware.Context) {
	var params models.TopicParams
	var topic models.Topic
	err := ctx.ReadJSON(&params)

	if err != nil {
		ctx.Go(1, "参数有误")
		return
	}

	user := ctx.Session().Get("user").(*models.User)

	topic.Title = params.Title
	topic.Content = params.Content
	topic.Shared = params.Shared
	topic.AuthorID= user.ID
	err = topic.Validate()

	if err != nil {
		fmt.Println("err", err)
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
func UpdateTopic(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")

	// 根据id获取topic
	var params models.TopicParams
	err := ctx.ReadJSON(&params)

	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	topic, err := TopicService.ByID(uint(id))

	if err != nil {
		ctx.Go(500, "话题不存在")
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	if topic.AuthorID != user.ID {
		ctx.Go(403, "禁止修改其他人的话题")
		return
	}

	err = TopicService.Update(topic, &params)
	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}

	ctx.Go(topic)
}

func TrashTopic(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")

	err := TopicService.DeleteByID(uint(id))
	if err != nil {
		ctx.Go(500, "删除失败")
		return
	}

	ctx.Go(id)
}

/**
 * 收藏
 */
func FavorTopic(ctx *middleware.Context) {
	var favor models.Favor
	err := ctx.ReadJSON(&favor)
	if err != nil {
		fmt.Println("error", err)
		ctx.Go(406, "参数有误")
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	favor.UserID = user.ID

	isFavor, err := TopicService.Favor(&favor)
	if err != nil {
		fmt.Println("err", err)
		ctx.Go(500, "操作失败")
		return
	}

	ctx.Go(isFavor)
}
