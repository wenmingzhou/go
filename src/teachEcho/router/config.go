package router

import (
	"github.com/labstack/echo"
	"io"
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
		return next(ctx)
	}
}
