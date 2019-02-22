package controllers

import (
	"groot/models"
	"groot/middleware"
)

type AnswerForm struct {
	TopicID int `json:"topicID"`
	AuthorID int `json:"authorID"`
}

func Comments(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = id

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
	topic.ID = id

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
	comt.Top = topic.Top
	comt.Nickname = user.Nickname
	comt.Avatar = user.Avatar
	err := comt.Save(topic)
	if err != nil {
		ctx.Error(500, "评论失败")
		return
	}

	topic.LastAvatar = user.Avatar
	topic.LastNickname = user.Nickname

	if (ctx.Client() != nil) {
		rt := make(map[string]interface{})
		rt["type"] = "comment"
		rt["data"] = map[string]interface{}{
			"topic": topic,
			"comment": comt,
		}

		go ctx.Client().Others(rt)
	}

	ctx.Go(comt)
}

/**
 * 回复
 */
func Reply(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = id

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
	reply.Top = topic.Top
	err := reply.Save(topic)
	if err != nil {
		ctx.Error(500, "回复失败")
		return
	}

	reply.ByID()
	topic.LastAvatar = user.Avatar
	topic.LastNickname = user.Nickname

	if (ctx.Client() != nil) {
		rt := make(map[string]interface{})
		rt["type"] = "reply"
		rt["data"] = map[string]interface{}{
			"topic": topic,
			"reply": reply,
		}
		go ctx.Client().Others(rt)
	}

	ctx.Go(reply)
}

// 评论作为答案
func AsAnswer(ctx *middleware.Context) {
	var form AnswerForm

	id, _ := ctx.Params().GetInt("id")
	ctx.ReadJSON(&form)
	topic := new(models.Topic)
	topic.ID = form.TopicID

	isExist := topic.IsExist()
	if !isExist {
		ctx.Error(404)
		return
	}

	if topic.AnswerID == id {
		id = 0
	}

	err := topic.UpdateField("answer_id", id)
	if err != nil {
		ctx.Error(500)
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	topic.Avatar = user.Avatar
	topic.Nickname = user.Nickname

	if (ctx.Client() != nil) {
		rt := make(map[string]interface{})
		rt["type"] = "answer"
		rt["data"] = map[string]interface{}{
			"topic": topic,
			"commentAuthorID": form.AuthorID,
		}
		go ctx.Client().Others(rt)
	}

	ctx.Go(topic)
}
