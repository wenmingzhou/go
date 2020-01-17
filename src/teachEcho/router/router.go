package router

import (
	"github.com/labstack/echo"
	"teachEcho/control"
)

var debug = true

func Run() {
	app := echo.New()

	app.Renderer = renderer

	app.Static("/static", "static")
	app.Static("/views", "views")
	app.GET("/", control.Index)

	app.GET("/login.html", control.LoginView)

	app.POST("/api/login", control.UserLogon)

	app.Start(":8021")
}
