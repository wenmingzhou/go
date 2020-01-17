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

	//路由分组
	admin := app.Group("/admin", ServerHeader)
	admin.GET("/index.html", control.AdminIndexView)

	app.POST("/api/login", control.UserLogon)
	app.Start(":8021")
}
