package controllers

import (
	"groot/models"
	"groot/middleware"
)

func Comments(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	comments, err := topic.GetComments()
	if err != nil {
		ctx.Go(500, "获取评论失败")
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
		ctx.Go(406, "帖子不存在")
		return
	}

	var comt models.Comment
	user := ctx.Session().Get("user").(*models.User)
	ctx.ReadJSON(&comt)
	if comt.Content == "" {
		ctx.Go(406, "回复内容不能为空")
		return
	}

	comt.AuthorID = user.ID
	err := comt.Save()
	if err != nil {
		ctx.Go(500, "评论失败")
		return
	}

	ctx.Go(comt)
}

/**
 * 回复
 */
func Reply(ctx *middleware.Context) {
	reply := new(models.Reply) 
	ctx.ReadJSON(reply)
	if reply.Content == "" {
		ctx.Go(406, "回复内容不能为空")
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	reply.AuthorID = user.ID
	err := reply.Save()
	if err != nil {
		ctx.Go(500, "回复失败")
		return
	}

	ctx.Go(reply)
}
