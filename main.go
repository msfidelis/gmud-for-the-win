package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createGMUD(c echo.Context) error {
	return c.JSON(http.StatusCreated, "olá")
}

func healthcheck(c echo.Context) error {
	return c.JSON(http.StatusCreated, "ok")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/gmud", createGMUD)
	e.GET("/", healthcheck)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
