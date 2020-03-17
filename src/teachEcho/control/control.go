package control

import (
	"github.com/labstack/echo"
)

func Index(ctx echo.Context) error {
	//fmt.Println("1111")
	//return ctx.String(200, "string")
	return ctx.Redirect(302, "login.html")
}

type pageLayUi struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
