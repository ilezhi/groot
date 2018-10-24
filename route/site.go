package Router

import (
	"github.com/kataras/iris"

	"groot/middleware"
	"groot/controllers"
)

func RegisterSite(app iris.Party) {
	// 获取页面topics
	app.Get("/topics/all", middleware.Handler(controllers.AllTopics))
	app.Get("/topics/awesome", middleware.Handler(controllers.AwesomeTopics))
	app.Get("/topics/department", middleware.Handler(controllers.DeptTopics))
	app.Get("/topics/my", middleware.Handler(controllers.MyTopics))
	app.Get("/topics/question", middleware.Handler(controllers.SolvedTopics))
	app.Get("/topics/answer", middleware.Handler(controllers.AnswerTopics))
	app.Get("/topics/favor/{id:int}", middleware.Handler(controllers.FavorTopics))
	app.Get("/topics/shared", middleware.Handler(controllers.SharedTopics))
	app.Get("/topics/tag/{id:int}", middleware.Handler(controllers.TagTopics))

	// 显示帖子详情
	app.Get("/topic/{id:int}", middleware.Handler(controllers.Topic))

	// 新增, 更新, 收藏, 点赞, 删除帖子
	app.Post("/topic/create", middleware.Handler(controllers.PublishTopic))
	app.Put("/topic/update/{id:int}", middleware.Handler(controllers.UpdateTopic))
	app.Post("/topic/favor/{id:int}", middleware.Handler(controllers.FavorTopic))
	app.Post("/topic/like/{id:int}", middleware.Handler(controllers.LikeTopic))
	app.Post("/topic/reply/{id:int}", middleware.Handler(controllers.Reply))
	
	app.Post("/tag/create", middleware.Handler(controllers.CreateTag))
	app.Post("/category/create", middleware.Handler(controllers.CreateCategory))

	// 获取帖子评论回复
	app.Get("/comments/{id:int}", middleware.Handler(controllers.Comments))
	app.Post("/comment/{id:int}", middleware.Handler(controllers.Comment))

	// 登录, 用户
	app.Post("/signin", middleware.Handler(controllers.SignIn))
	app.Get("/user/info", middleware.Handler(controllers.UserInfo))	
}
