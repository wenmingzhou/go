package router

import (
	"github.com/labstack/echo"
	"tech/control"
)

func Run(){
	app :=echo.New()
	app.GET("/",control.Index)
	app.Start(":8080")
}
