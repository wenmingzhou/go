package ginCode

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
	c.HTML(http.StatusOK, "home.html", nil)
}


func vipHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "vip.html", nil)
}