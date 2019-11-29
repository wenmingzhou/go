package main

import (
	"net/http"
	"news/control"
)

func main() {
	fsh := http.FileServer(http.Dir("./src/news/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))

	http.HandleFunc("/", control.IndexView)
	http.HandleFunc("/api/index/data/", control.IndexData)

	//列表页面
	http.HandleFunc("/list", control.ListView)
	http.HandleFunc("/api/list/data/", control.ListData)

	//添加页面
	http.HandleFunc("/add", control.ViewArticleAdd)
	http.HandleFunc("/api/article/add/", control.ApiArticleAdd)

	http.ListenAndServe(":8080", nil)
}
