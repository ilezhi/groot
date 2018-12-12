package Router

import (
	"github.com/kataras/iris"

	"groot/middleware"
	"groot/controllers"
)

func Register(app *iris.Application) {
	user := app.Party("/ajax/v1", middleware.Handler(middleware.IsLogin))
	{
		// 获取页面topics
		user.Get("/topics/top", middleware.Handler(controllers.Top))
		user.Get("/topics/all", middleware.Handler(controllers.AllTopics))
		user.Get("/topics/awesome", middleware.Handler(controllers.AwesomeTopics))
		user.Get("/topics/department", middleware.Handler(controllers.DeptTopics))
		user.Get("/topics/my", middleware.Handler(controllers.MyTopics))
		user.Get("/topics/question", middleware.Handler(controllers.SolvedTopics))
		user.Get("/topics/answer", middleware.Handler(controllers.AnswerTopics))
		user.Get("/topics/favor/{id:int}", middleware.Handler(controllers.FavorTopics))
		user.Get("/topics/shared", middleware.Handler(controllers.SharedTopics))
		user.Get("/topics/tag/{id:int}", middleware.Handler(controllers.TagTopics))

		// 显示帖子详情
		user.Get("/topic/{id:int}", middleware.Handler(controllers.Topic))

		// 新增, 更新, 收藏, 点赞, 删除帖子
		user.Post("/topic/create", middleware.Handler(controllers.PublishTopic))
		user.Put("/topic/update/{id:int}", middleware.Handler(controllers.UpdateTopic))
		user.Post("/topic/favor/{id:int}", middleware.Handler(controllers.FavorTopic))
		user.Post("/like/{id:int}", middleware.Handler(controllers.Like))
		// user.Post("/topic/reply/{id:int}", middleware.Handler(controllers.Reply))

		user.Post("/tag/create", middleware.Handler(controllers.CreateTag))
		user.Post("/category/create", middleware.Handler(controllers.CreateCategory))

		// 获取帖子评论回复
		user.Get("/comments/{id:int}", middleware.Handler(controllers.Comments))
		user.Post("/comment/{id:int}", middleware.Handler(controllers.Comment))
		user.Post("/comment/reply/{id:int}", middleware.Handler(controllers.Reply))

		// 登录, 用户
		user.Get("/user/info", middleware.Handler(controllers.UserInfo))
		user.Get("/user", middleware.Handler(controllers.LoginUser))
	}

	admin := app.Party("/ajax/v1", middleware.Handler(middleware.IsAdmin))
	{
		admin.Put("/topic/top/{id:int}", middleware.Handler(controllers.SetTop))
		admin.Put("/topic/awesome/{id:int}", middleware.Handler(controllers.SetAwesome))
		admin.Delete("/topic/delete/{id:int}", middleware.Handler(controllers.TrashTopic))
	}
	
	site := app.Party("/ajax/v1/")
	{
		site.Post("signin", middleware.Handler(controllers.SignIn))
		site.Post("signup", middleware.Handler(controllers.SignUp))
		site.Get("exist", middleware.Handler(controllers.IsExistByEmail))
		site.Get("departments", middleware.Handler(controllers.Departments))
	}
}
