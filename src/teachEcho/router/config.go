package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"io"
	"teachEcho/model"
	"text/template"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	/*if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}*/
	if debug {
		t.templates =
			template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

//第一次运行程序加载  加载到内存中
var renderer = &TemplateRenderer{
	templates: template.Must(template.ParseFiles("./views/login.html", "./views/index.html")),
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "Echo/999")
		tokenString := ctx.FormValue("token")
		claims := model.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("house365"), nil
		})

		if err == nil && token.Valid {
			ctx.Set("uid", claims.Id)
			//验证通过
			return next(ctx)
		} else {
			return ctx.JSON(utils.ErrJwt("jwt token 验证失败"))
		}

	}
}
