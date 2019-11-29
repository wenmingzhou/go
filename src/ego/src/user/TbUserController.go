package user

import (
	"ego/src/commons"
	"encoding/json"
	"net/http"
)

func UserHandler()  {
	commons.Router.HandleFunc("/login",loginController)
}
//登陆
func loginController(w http.ResponseWriter,r *http.Request)  {
	username :=r.FormValue("username")
	password :=r.FormValue("password")
	er :=loginService(username,password)
	//把结构体转换为json数据
	b,_ :=json.Marshal(er)
	//设置响应内容为json
	w.Header().Set("Content-type","application/json;charset=utf-8")
	w.Write(b)
}
