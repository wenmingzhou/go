package control

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"teachEcho/model"
	"time"
)

type login struct {
	Num  string `json:"num"`
	Pass string `json:"pass"`
}

func UserLogon(ctx echo.Context) error {
	ipt := login{}
	err := ctx.Bind(&ipt) // 把传输的数据绑定到结构体中
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.UserLogin(ipt.Num)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("用户名错误", err.Error()))
	}
	if mod.Pass != ipt.Pass {
		return ctx.JSON(utils.ErrIpt("密码错误"))
	}

	// Create the Claims
	claims := model.UserClaims{
		Id:   mod.Id,
		Num:  mod.Num,
		Name: mod.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("house365"))
	fmt.Printf("%v %v", ss, err)

	return ctx.JSON(utils.Succ("登录成功", ss))

}
