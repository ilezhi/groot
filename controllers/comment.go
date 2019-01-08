package controllers

import (
	"groot/models"
	"groot/middleware"
)

func Comments(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	user := ctx.Session().Get("user").(*models.User)
	comments, err := topic.GetComments(user.ID)
	if err != nil {
		ctx.Error(500, "获取评论失败")
		return
	}

	ctx.Go(comments)
}

/**
 * 评论
 */
func Comment(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	exist := topic.IsExist()
	if !exist {
		ctx.Error(404)
		return
	}

	var comt models.Comment
	user := ctx.Session().Get("user").(*models.User)
	ctx.ReadJSON(&comt)
	if comt.Content == "" {
		ctx.Error(422, "回复内容不能为空")
		return
	}

	comt.AuthorID = user.ID
	comt.TopicID = topic.ID
	comt.Nickname = user.Nickname
	comt.Avatar = user.Avatar
	err := comt.Save(topic)
	if err != nil {
		ctx.Error(500, "评论失败")
		return
	}

	topic.LastAvatar = user.Avatar
	topic.LastNickname = user.Nickname

	rt := make(map[string]interface{})
	rt["type"] = "comment"
	rt["data"] = map[string]interface{}{
		"topic": topic,
		"comment": comt,
	}
	go ctx.Client().Others(rt)
	ctx.Go(comt)
}

/**
 * 回复
 */
func Reply(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	exist := topic.IsExist()
	if !exist {
		ctx.Error(404, "帖子不存在")
	}

	reply := new(models.Reply)
	ctx.ReadJSON(reply)
	if reply.Content == "" {
		ctx.Error(422, "回复内容不能为空")
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	reply.AuthorID = user.ID
	reply.TopicID = topic.ID
	err := reply.Save(topic)
	if err != nil {
		ctx.Error(500, "回复失败")
		return
	}

	reply.ByID()
	topic.LastAvatar = user.Avatar
	topic.LastNickname = user.Nickname

	rt := make(map[string]interface{})
	rt["type"] = "reply"
	rt["data"] = map[string]interface{}{
		"topic": topic,
		"reply": reply,
	}
	go ctx.Client().Others(rt)
	ctx.Go(reply)
}
