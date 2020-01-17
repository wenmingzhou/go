package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//user 表
type User struct {
	Id     int       `json:"id"`
	Num    string    `json:"num"` //账号
	Name   string    `json:"name"`
	Pass   string    `json:"pass"`
	Phone  string    `json:"phone"`
	Email  string    `json:"email"`
	Status int       `json:"status"`
	Ctime  time.Time `json:"ctime"`
}

//用户登录
func UserLogin(num string) (User, error) {
	mod := User{}
	err := DB.Get(&mod, "select * from user where num=? limit 1", num)
	return mod, err
}

//token 要携带的数据
type UserClaims struct {
	Id   int    `json:"id"`
	Num  string `json:"num"`
	Name string `json:"name"`
	jwt.StandardClaims
}
