package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"

	"groot/route"
)

func main() {
	app := iris.New()
	
	app.Use(logger.New())
	
	Router.Register(app)

	app.Run(
		iris.Addr("localhost:9000"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
