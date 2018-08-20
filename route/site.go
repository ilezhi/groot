package Router

import (
	"github.com/kataras/iris"

	"groot/controllers"
)

func RegisterSite(app iris.Party) {
	app.Get("/topic", topic.List)
	app.Get("/topic/{id:int}", topic.ByID)
	app.Post("/topic/create", topic.Create)
}
