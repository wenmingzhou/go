package main

import (
	"net/http"
	"tech/control"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", control.IndexView)
	http.HandleFunc("/api/index/data", control.IndexData)

	http.HandleFunc("/detail", control.DetailView)

	http.HandleFunc("/edit", control.EditView)
	http.HandleFunc("/api/article/edit", control.ApiArticleEdit)

	http.HandleFunc("/api/article/page", control.ApiListPage)

	http.HandleFunc("/list", control.ListView)
	http.HandleFunc("/api/list/data", control.ListData)

	http.HandleFunc("/api/list/del", control.ListDel)

	http.HandleFunc("/add", control.ArticleAddView)
	http.HandleFunc("/api/article/add", control.ApiArticleAdd)

	http.HandleFunc("/api/upload", control.ApiUpload)

	http.ListenAndServe(":8090", nil)
}
