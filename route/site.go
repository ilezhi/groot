package Router

import (
	"github.com/kataras/iris"

	"groot/controllers"

	"groot/tools"
)

func RegisterSite(app iris.Party) {
	app.Get("/topics", tools.Handler(controllers.TopicList))
	app.Get("/topic/{id:int}", tools.Handler(controllers.TopicByID))
	app.Post("/topic/create", tools.Handler(controllers.CreateTopic))

	app.Get("/tags", tools.Handler(controllers.TagList))
	app.Get("/tag/{id:int}", tools.Handler(controllers.TagByID))
	app.Post("/tag/create", tools.Handler(controllers.CreateTag))
}
