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

func f1(w http.ResponseWriter, r *http.Request) {
	//2解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles failed,err:%v", err)
		return
	}
	//3渲染模板
	name := "小王子"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tempDemo", f1)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("HTTP server start failed,err:%v", err)
		return
	}
}
