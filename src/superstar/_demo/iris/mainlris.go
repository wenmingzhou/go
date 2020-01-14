package main

import "github.com/kataras/iris"

func main()  {
	app :=iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello world --from iris")
	})
	htmlEngine :=iris.HTML("./_demo/iris/",".html")
	app.RegisterView(htmlEngine)
	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title","测试页面")
		ctx.ViewData("Content","Hell world-- from template")
		ctx.View("hello.html")
	})
	app.Run(iris.Addr(":8081"),iris.WithCharset("UTF-8"))
}