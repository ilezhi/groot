package Router

import (
	"github.com/kataras/iris"

	"groot/controllers"

	"groot/tools"
)

func RegisterSite(app iris.Party) {
	app.Get("/topics", tools.Handler(controllers.Topics))
	app.Get("/topic/{id:int}", tools.Handler(controllers.Topic))
	app.Post("/topic/create", tools.Handler(controllers.CreateTopic))
	app.Post("/topic/update/{id:int}", tools.Handler(controllers.UpdateTopic))
	app.Post("/topic/awesome/{id:int}", tools.Handler(controllers.AwesomeTopic))

	app.Get("/tags", tools.Handler(controllers.TagList))
	app.Get("/tag/{id:int}", tools.Handler(controllers.TagByID))
	app.Post("/tag/create", tools.Handler(controllers.CreateTag))
}
