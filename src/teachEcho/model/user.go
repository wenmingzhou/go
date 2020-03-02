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

//返回分页的数据
func UserPage(pi, ps int) ([]User, error) {
	mods := make([]User, 0, ps)
	err := DB.Select(&mods, "select * from user limit ?, ?", (pi-1)*ps, ps)
	return mods, err
}

//返回当前条件的总数
func UserCount() int {
	count := 0
	DB.Get(&count, "select count(*) as count from user")
	return count
}
