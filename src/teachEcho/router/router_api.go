package router

import (
	"github.com/labstack/echo"
	"teachEcho/control"
)

func ApiRouter(api *echo.Group) {
	api.POST("/login", control.UserLogin)
	api.GET("/class/all", control.ClassAll)
	api.GET("/class/page", control.ClassPage)
	api.GET("/class/get/:id", control.ClassGet)
}
