package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		toPath := c.DefaultQuery("next", "/index")
		var u UserInfo
		err := c.ShouldBind(&u)
		fmt.Println(u.Username)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "绑定失败",
			})
			return
		}
		if u.Username == "" || u.Password == "" {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码不能为空",
			})
			return
		}
		if u.Username == "guan" && u.Password == "123" {
			//登录成功
			//1 设置cookie
			c.SetCookie("username", u.Username, 120, "/", "127.0.0.1", false, true)
			//跳转到index页面
			c.Redirect(302, toPath)
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码错误",
			})
			return
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}

}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func homeHandler(c *gin.Context) {
	//在返回页面之前要校验是否存在username的cookie
	//获取cookie
	username, err := c.Cookie("username")
	if err != nil {
		c.Redirect(302, "/login")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})
}

//基于Cookie实现用户登录论证的中间件
func cookieMiddleware(c *gin.Context) {
	//在返回页面之前要校验是否存在username的cookie
	//获取cookie

	toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
	username, err := c.Cookie("username")
	if err != nil {
		c.Redirect(302, toPath)
		return
	}
	//用户已经登录了
	c.Set("username", username) //在上下文中设置一个键值对
	c.Next()                    //继续后续的处理函数
}

func vipHandler(c *gin.Context) {
	tmpUsername, ok := c.Get("username")
	if !ok {
		//如果取不到值，说明前面的中间件出问题了
		c.Redirect(302, "/index")
		return
	}
	username, ok := tmpUsername.(string)
	if !ok {
		//类型断言失败
		c.Redirect(302, "/index")
		return
	}
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
	})
}

//cookie示例
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", cookieMiddleware, vipHandler)
	r.Run(":9090")
}
