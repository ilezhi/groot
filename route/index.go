package Router

import (
	"github.com/kataras/iris"
)

func Register(app *iris.Application) {
	RegisterSite(app.Party("/ajax"))
	RegisterAdmin(app.Party("/admin"))
}
