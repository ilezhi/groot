package Router

import (
	"github.com/kataras/iris"

	"groot/controllers"

	"groot/tools"
)

func RegisterSite(app iris.Party) {
	app.Get("/topics/all", tools.Handler(controllers.Topics))
	app.Get("/topics/awesome", tools.Handler(controllers.AwesomeTopics))
	// app.Get("/topics/department", tools.Handler(controllers.DetpTopics))
	app.Get("/topic/{id:int}", tools.Handler(controllers.Topic))

	app.Post("/topic/create", tools.Handler(controllers.CreateTopic))
	app.Post("/topic/update/{id:int}", tools.Handler(controllers.UpdateTopic))
	app.Post("/topic/awesome/{id:int}", tools.Handler(controllers.AwesomeTopic))

	app.Get("/tags", tools.Handler(controllers.Tags))
	app.Get("/tag/{id:int}", tools.Handler(controllers.Tag))
	app.Post("/tag/create", tools.Handler(controllers.CreateTag))

	// 搜索相关
	app.Get("/search/tag/{name:string}", tools.Handler(controllers.SearchTag))
}
