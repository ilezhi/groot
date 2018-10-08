package controllers

import (
	// "math/rand"
	// "time"
	"fmt"
	"groot/models"
	"groot/middleware"
)

type TopicParams struct {
	Title 	string		`json:"title"`
	Content string		`json:"content"`
	Tags 		[]uint		`json:"tags"`
	Shared 	bool			`json:"shared"`
}

func AllTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)

	topic.UpdatedAt = lastID

	topics, err := topic.All()
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func AwesomeTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)

	topic.UpdatedAt = lastID

	topics, err := topic.Awesomes()
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func DeptTopics(ctx *middleware.Context) {
	lastID, err := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)

	topic.UpdatedAt = lastID
	topic.AuthorID = user.ID

	topics, err := topic.Department(user.DeptID)
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func MyTopics(ctx *middleware.Context) {
	lastID, err := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)

	topic.UpdatedAt = lastID
	topic.AuthorID = user.ID

	topics, err := topic.UnSolved()
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func SolvedTopics(ctx *middleware.Context) {
	lastID, err := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)

	topic.UpdatedAt = lastID
	topic.AuthorID = user.ID

	topics, err := topic.Solved()
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func AnswerTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)

	topic.UpdatedAt = lastID
	topic.AuthorID = user.ID

	topics, err := topic.CommentAsAnswer()
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func FavorTopics(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	lastID, _ := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)
	
	topic.UpdatedAt = lastID

	topics, err := topic.GetByCategory(uint(id))
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func SharedTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)

	topic.UpdatedAt = lastID
	topic.AuthorID = user.ID

	topics, err := topic.SharedList()
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func TagTopics(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	lastID, _ := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)

	topic.UpdatedAt = lastID
	topic.AuthorID = user.ID

	topics, err := topic.GetByTag(uint(id))
	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(topics)
}

func Topic(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	isExist := topic.IsExist()
	if !isExist {
		ctx.Go(404, "帖子不翼而飞")
		return
	}

	err := topic.GetTags()
	if err != nil {
		ctx.Go(500, "获取帖子标签失败")
		return
	}

	ctx.Go(topic)
}

/**
 * 发布话题
 * content, title, tags, shared
 */
func PublishTopic(ctx *middleware.Context) {
	var params TopicParams
	err := ctx.ReadJSON(&params)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)
	topic.Title = params.Title
	topic.Content = params.Content
	topic.Shared = params.Shared
	topic.AuthorID= user.ID

	err = topic.Validate()
	if err != nil {
		fmt.Println("err", err)
		ctx.Go(406, "参数有误")
		return
	}

	// 添加话题和tag
	err = topic.Insert(&params.Tags)
	if err != nil {
		ctx.Go(500, "新增失败")
		return
	}

	ctx.Go(topic)
}

/**
 * 更新
 */
func UpdateTopic(ctx *middleware.Context) {
	var params TopicParams
	err := ctx.ReadJSON(&params)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	if params.Content == "" {
		ctx.Go(406, "帖子内容不能为空")
		return
	}

	if len(params.Tags) < 1 {
		ctx.Go(406, "帖子至少需要添加一个标签")
		return
	}

	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	isExist := topic.IsExist()
	if !isExist {
		ctx.Go(404, "帖子不存在")
		return
	}

	topic.Content = params.Content
	topic.Shared = params.Shared
	topic.Update(&params.Tags)

	ctx.Go(topic)
}

func TrashTopic(ctx *middleware.Context) {
	ctx.Go(200)
}

/**
 * 收藏
 */
func FavorTopic(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	isExist := topic.IsExist()
	if !isExist {
		ctx.Go(404, "帖子不存在")
		return
	}

	var favor models.Favor
	err := ctx.ReadJSON(&favor)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	favor.TopicID = uint(id)
	favor.UserID = user.ID
	isFavor := favor.IsFavor()
	if isFavor {
		// 取消收藏
		err = favor.Delete()
		if err != nil {
			ctx.Go(500, "取消收藏失败")
			return
		}
	} else {
		// 收藏
		err = favor.Insert()
		if err != nil {
			ctx.Go(500, "收藏失败")
			return
		}
	}	
		
	ctx.Go(id)
}

/**
 * 点赞
 */
func LikeTopic(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	like := new(models.Like)
	user := ctx.Session().Get("user").(*models.User)
	like.TargetID = uint(id)
	like.UserID = user.ID
	like.Type = "topic"
	
	isExist := like.IsExist()
	var err error
	if isExist {
		// 取消点赞
		err = like.Delete()

		if err != nil {
			ctx.Go(500, "取消点赞失败")
			return
		}
	} else {
		err = like.Save()
		if err != nil {
			ctx.Go(500, "点赞失败")
			return
		}
	}

	ctx.Go(id)
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
