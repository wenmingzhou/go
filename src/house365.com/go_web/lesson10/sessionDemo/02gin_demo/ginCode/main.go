package ginCode

import (
	"github.com/gin-gonic/gin"
	"ginsession"
)
//测试Session服务的gin demo
func main()  {
	r :=gin.Default()
	r.LoadHTMLGlob("templates/*")
	//session中间件应该作为一个全局的中间件
	//初始化全局的MgrObj对象
	ginsession.InitMgr()  //后面可以扩展redis版
	r.Use(ginsession.SessionMiddleware(ginsession.MgrObj))
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", vipHandler)
	r.Run()
}
