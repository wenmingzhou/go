package control

import "github.com/labstack/echo"

func LoginView(ctx echo.Context) error {
	return ctx.Render(200, "login.html", nil)
}
