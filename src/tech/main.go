package main

import (
	"net/http"
	"tech/control"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", control.IndexView)
	http.HandleFunc("/api/index/data", control.IndexData)

	http.HandleFunc("/list", control.ListView)
	http.HandleFunc("/api/list/data", control.ListData)

	http.HandleFunc("/api/list/del", control.ListDel)

	http.HandleFunc("/add", control.ArticleAddView)
	http.HandleFunc("/api/article/add", control.ApiArticleAdd)

	http.ListenAndServe(":8090", nil)
}
