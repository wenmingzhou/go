package control

import "github.com/labstack/echo"

func Index(ctx echo.Context) error {
	return ctx.String(200, "string")
}

type pageLayUi struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
