package Router

import (
	"github.com/kataras/iris"

	"groot/controllers"

	"groot/middleware"
)

func RegisterSite(app iris.Party) {
	app.Get("/topics/all", middleware.Handler(controllers.Topics))
	app.Get("/topics/awesome", middleware.Handler(controllers.AwesomeTopics))
	// app.Get("/topics/department", middleware.Handler(controllers.DetpTopics))
	app.Get("/topics/my", middleware.Handler(controllers.MyTopics))
	app.Get("/topics/question", middleware.Handler(controllers.QTopics))
	app.Get("/topics/answer", middleware.Handler(controllers.ATopics))
	app.Get("/topic/{id:int}", middleware.Handler(controllers.Topic))

	app.Post("/topic/create", middleware.Handler(controllers.CreateTopic))
	app.Post("/topic/update/{id:int}", middleware.Handler(controllers.UpdateTopic))
	// 收藏
	app.Post("/topic/favor", middleware.Handler(controllers.FavorTopic))

	app.Get("/tags", middleware.Handler(controllers.Tags))
	app.Get("/tag/{id:int}", middleware.Handler(controllers.Tag))
	app.Post("/tag/create", middleware.Handler(controllers.CreateTag))

	// 评论
	app.Post("/topic/{id:int}/comment", middleware.Handler(controllers.Comment))

	// 搜索相关
	app.Get("/search/tag/{name:string}", middleware.Handler(controllers.SearchTag))

	// 登录
	app.Get("/signin", middleware.Handler(controllers.Login))
	
	// 用户相关
	app.Post("/user/create", middleware.Handler(controllers.CreateUser))
	app.Post("/category/create", middleware.Handler(controllers.CreateCategory))
}
