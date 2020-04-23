package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed,err:%v", err)
		return
	}
	//3渲染模板
	u := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}

	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	m := map[string]interface{}{
		"name":   "小公主",
		"gender": "女",
		"age":    16,
	}
	//传递多个变量
	err = t.Execute(w, map[string]interface{}{
		"u":     u,
		"m":     m,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("render template failed,err:%v", err)
		return
	}

}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("HTTP server start failed,err:%v", err)
		return
	}
}
