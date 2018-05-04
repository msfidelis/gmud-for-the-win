package main

import (
	"github.com/kataras/iris"
)

var addr = iris.Addr("0.0.0.0:3000")

func main() {
	app := iris.New()
	app.Get("/healthcheck", index)
}

func healthcheck(ctx iris.Context) {
	ctx.JSON(iris.Map{"health": "status"})
}
