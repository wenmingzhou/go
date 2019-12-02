package main

import (
	"ego/src/commons"
	"ego/src/item"
	"ego/src/item/cat"
	"ego/src/user"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

// main.go

func welcome(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("view/login.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, nil)
}

//restful显示页面
func showPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmpl, _ := template.ParseFiles("view/" + vars["page"] + ".html")
	tmpl.Execute(w, nil)
}
func main() {

	/*http.HandleFunc("/", welcome)
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	//调用所用user模块的handler
	user.UserHandler()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}*/
	commons.Router.HandleFunc("/", welcome)
	//满足/page/{page}格式的处理
	commons.Router.HandleFunc("/page/{page}", showPage)
	commons.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	user.UserHandler()   //调用所用user模块的handler
	item.ItemHandler()   //商品
	cat.ItemCatHandler() //分类
	http.ListenAndServe(":8888", commons.Router)

}
