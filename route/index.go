package Router

import (
	"github.com/kataras/iris"
)

func Register(app *iris.Application) {
	RegisterSite(app.Party("/api"))
	RegisterAdmin(app.Party("/admin"))
}
