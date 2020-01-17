package router

import (
	"github.com/labstack/echo"
	"teachEcho/control"
)

func Run() {
	app := echo.New()
	app.Static("/static", "static")
	app.Static("/views", "views")
	app.GET("/", control.Index)
	app.POST("/api/login", control.UserLogon)

	app.Start(":8021")
}
