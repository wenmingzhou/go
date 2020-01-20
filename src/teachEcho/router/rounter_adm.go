package router

import (
	"github.com/labstack/echo"
	"teachEcho/control"
)

//必须要token
func AdmRouter(adm *echo.Group) {
	adm.POST("/class/add", control.ClassAdd)
	adm.GET("/class/drop/:id", control.ClassDrop)

	adm.POST("/class/edit", control.ClassEdit)
}
