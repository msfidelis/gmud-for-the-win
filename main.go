package main

import (
	"github.com/kataras/iris"
)

var addr = iris.Addr("0.0.0.0:3000")

func main() {
	app := iris.New()
	app.Get("/healthcheck", healthcheck)
	app.Post("/gmud", gmud)
	app.Run(addr)
}

func healthcheck(ctx iris.Context) {
	ctx.JSON(iris.Map{"status": 200})
}

func gmud(ctx iris.Context) {
	ctx.JSON(iris.Map{"feature": "create GMUD PDF"})
}
