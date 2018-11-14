package controllers

import (
	// "math/rand"
	// "time"
	// "fmt"
	"groot/models"
	"groot/middleware"
)

type TopicForm struct {
	ID			 		uint      `json:"id,omitempty"`
	Title 	 		string		`json:"title"`
	Content  		string		`json:"content,omitempty"`
	Tags 		 		[]uint		`json:"tags,omitempty"`
	Shared 	 		bool			`json:"shared"`
	Type     		string    `json:"type,omitempty"`
	AuthorID 		uint			`json:"authorID,omitempty"`
	IsLike   		bool      `json:"isLike"`
	IsFavor  		bool      `json:"isFavor"`
	CategoryID 	uint			`json:"categoryID,omitempty"`
	Nickname    string    `json:"nickname,omitempty"`
	Avatar			string		`json:"avatar,omitempty"`
}

func AllTopics(ctx *middleware.Context) {
	lastID, _ := ctx.URLParamInt64("lastID")
	topic := new(models.Topic)

	topic.ActiveAt = lastID

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

	topic.ActiveAt = lastID

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

	topic.ActiveAt = lastID
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

	topic.ActiveAt = lastID
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

	topic.ActiveAt = lastID
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

	topic.ActiveAt = lastID
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
	
	topic.ActiveAt = lastID

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

	topic.ActiveAt = lastID
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

	topic.ActiveAt = lastID
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

	err := topic.FindByID()
	if err != nil {
		ctx.Go(500, "获取帖子标签失败")
		return
	}

	topic.View += 1
	topic.UpdateView()
	
	user := ctx.Session().Get("user").(*models.User)
	topic.GetCount(user.ID)
	topic.IsFull = true

	ctx.Go(topic)
}

/**
 * 发布话题
 * content, title, tags, shared
 */
func PublishTopic(ctx *middleware.Context) {
	var form TopicForm
	err := ctx.ReadJSON(&form)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	topic := new(models.Topic)
	user := ctx.Session().Get("user").(*models.User)
	topic.Title = form.Title
	topic.Content = form.Content
	topic.Shared = form.Shared
	topic.AuthorID= user.ID
	topic.Issue = true

	err = topic.Validate()
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	// 添加话题和tag
	err = topic.Save(&form.Tags)
	if err != nil {
		ctx.Go(500, "新增失败")
		return
	}

	topic.GetTags()

	topic.Nickname = user.Nickname
	topic.Avatar = user.Avatar
	rt := make(map[string]interface{})
	rt["type"] = "topic"
	rt["data"] = topic
	go ctx.Client().Others(rt)
	ctx.Go(topic)
}

/**
 * 更新
 */
func UpdateTopic(ctx *middleware.Context) {
	var form TopicForm
	err := ctx.ReadJSON(&form)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	if form.Content == "" {
		ctx.Go(406, "帖子内容不能为空")
		return
	}

	if len(form.Tags) < 1 {
		ctx.Go(406, "帖子至少需要添加一个标签")
		return
	}

	id, _ := ctx.Params().GetInt("id")
	user := ctx.Session().Get("user").(*models.User)
	topic := new(models.Topic)
	topic.ID = uint(id)

	isExist := topic.IsExist()
	if !isExist {
		ctx.Go(404, "帖子不存在")
		return
	}

	topic.Content = form.Content
	topic.Shared = form.Shared
	err = topic.Update(&form.Tags)
	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}

	topic.GetTags()
	topic.Nickname = user.Nickname
	topic.Avatar = user.Avatar

	rt := make(map[string]interface{})
	rt["type"] = "topic"
	rt["action"] = "update"
	rt["data"] = topic
	go ctx.Client().Others(rt)
	ctx.Go(topic)
}

func TrashTopic(ctx *middleware.Context) {
	ctx.Go(200)
}

/**
 * 收藏
 */
func FavorTopic(ctx *middleware.Context) {
	var form TopicForm
	id, _ := ctx.Params().GetInt("id")
	
	topic := new(models.Topic)
	topic.ID = uint(id)
	
	isExist := topic.IsExist()
	if !isExist {
		ctx.Go(404, "帖子不存在")
		return
	}
	
	favor := new(models.Favor)
	err := ctx.ReadJSON(&form)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	user := ctx.Session().Get("user").(*models.User)
	favor.TopicID = uint(id)
	favor.UserID = user.ID
	favor.CategoryID = form.CategoryID
	isFavor := favor.IsExist()
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

	form.IsFavor = !isFavor
	form.ID = uint(id)
	form.Nickname = user.Nickname
	form.Avatar = user.Avatar
	rt := make(map[string]interface{})
	rt["type"] = "favor"
	rt["data"] = form
	go ctx.Client().Others(rt)

	ctx.Go(form.IsFavor)
}

/**
 * 点赞
 */
func Like(ctx *middleware.Context) {
	var form TopicForm
	id, _ := ctx.Params().GetInt("id")
	user := ctx.Session().Get("user").(*models.User)
	like := new(models.Like)
	ctx.ReadJSON(&form)

	like.TargetID = uint(id)
	like.UserID = user.ID
	like.Type = form.Type

	isLike := like.IsExist()
	var err error
	if isLike {
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

	form.IsLike = !isLike
	form.ID = uint(id)
	form.Nickname = user.Nickname
	form.Avatar = user.Avatar
	rt := make(map[string]interface{})
	rt["type"] = "like"
	rt["data"] = form
	go ctx.Client().Others(rt)

	ctx.Go(form.IsLike)
}

func SetTop(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	isExist := topic.IsExist()

	if !isExist {
		ctx.Go(404, "话题不存在")
		return
	}

	topic.Top = !topic.Top
	err := topic.UpdateField("top", topic.Top)
	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}

	err = topic.FindFullByID()
	if err != nil {
		ctx.Go(500, "获取帖子失败")
		return
	}

	rt := make(map[string]interface{})
	rt["type"] = "top"
	rt["data"] = topic
	go ctx.Client().Others(rt)

	ctx.Go(topic)
}

func SetAwesome(ctx *middleware.Context) {
	id, _ := ctx.Params().GetInt("id")
	topic := new(models.Topic)
	topic.ID = uint(id)

	isExist := topic.IsExist()

	if !isExist {
		ctx.Go(404, "话题不存在")
		return
	}

	topic.Awesome = !topic.Awesome
	err := topic.UpdateField("awesome", topic.Awesome)
	if err != nil {
		ctx.Go(500, "更新失败")
		return
	}

	err = topic.FindFullByID()
	if err != nil {
		ctx.Go(500, "获取帖子失败")
		return
	}

	rt := make(map[string]interface{})
	rt["type"] = "awesome"
	rt["data"] = topic
	go ctx.Client().Others(rt)

	ctx.Go(topic)
}
