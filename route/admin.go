package Router

import (
	"github.com/kataras/iris"
)

func RegisterAdmin(app iris.Party) {
	app.Get("/json", func (ctx iris.Context) {
		ctx.JSON(iris.Map{"name": "admin route"})
	})
}
