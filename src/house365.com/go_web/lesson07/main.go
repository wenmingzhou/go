package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//2解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles failed,err:%v", err)
		return
	}
	//3渲染模板
	msg := "小王子"
	t.ExecuteTemplate(w, "index.tmpl", msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles failed,err:%v", err)
		return
	}
	//3渲染模板
	msg := "小王子"
	t.ExecuteTemplate(w, "home.tmpl", msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
}
