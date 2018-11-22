package Router

import (
	"github.com/kataras/iris"

	"groot/middleware"
	"groot/controllers"
)

func Register(app *iris.Application) {
	site := app.Party("/ajax/v1", middleware.Handler(middleware.IsLogin))
	{
		// 获取页面topics
		site.Get("/topics/top", middleware.Handler(controllers.Top))
		site.Get("/topics/all", middleware.Handler(controllers.AllTopics))
		site.Get("/topics/awesome", middleware.Handler(controllers.AwesomeTopics))
		site.Get("/topics/department", middleware.Handler(controllers.DeptTopics))
		site.Get("/topics/my", middleware.Handler(controllers.MyTopics))
		site.Get("/topics/question", middleware.Handler(controllers.SolvedTopics))
		site.Get("/topics/answer", middleware.Handler(controllers.AnswerTopics))
		site.Get("/topics/favor/{id:int}", middleware.Handler(controllers.FavorTopics))
		site.Get("/topics/shared", middleware.Handler(controllers.SharedTopics))
		site.Get("/topics/tag/{id:int}", middleware.Handler(controllers.TagTopics))

		// 显示帖子详情
		site.Get("/topic/{id:int}", middleware.Handler(controllers.Topic))

		// 新增, 更新, 收藏, 点赞, 删除帖子
		site.Post("/topic/create", middleware.Handler(controllers.PublishTopic))
		site.Put("/topic/update/{id:int}", middleware.Handler(controllers.UpdateTopic))
		site.Post("/topic/favor/{id:int}", middleware.Handler(controllers.FavorTopic))
		site.Post("/like/{id:int}", middleware.Handler(controllers.Like))
		// site.Post("/topic/reply/{id:int}", middleware.Handler(controllers.Reply))

		site.Post("/tag/create", middleware.Handler(controllers.CreateTag))
		site.Post("/category/create", middleware.Handler(controllers.CreateCategory))

		// 获取帖子评论回复
		site.Get("/comments/{id:int}", middleware.Handler(controllers.Comments))
		site.Post("/comment/{id:int}", middleware.Handler(controllers.Comment))
		site.Post("/comment/reply/{id:int}", middleware.Handler(controllers.Reply))

		// 登录, 用户
		site.Get("/user/info", middleware.Handler(controllers.UserInfo))
		site.Get("/user", middleware.Handler(controllers.LoginUser))
	}

	admin := app.Party("/ajax/v1", middleware.Handler(middleware.IsAdmin))
	{
		admin.Put("/topic/top/{id:int}", middleware.Handler(controllers.SetTop))
		admin.Put("/topic/awesome/{id:int}", middleware.Handler(controllers.SetAwesome))
		admin.Delete("/topic/delete/{id:int}", middleware.Handler(controllers.TrashTopic))
	}
	
	app.Post("/ajax/v1/signin", middleware.Handler(controllers.SignIn))
}
