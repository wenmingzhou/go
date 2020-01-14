package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"tech/model"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", IndexView)
	http.HandleFunc("/api/index/data", IndexData)

	http.HandleFunc("/list", ListView)
	http.HandleFunc("/api/list/data", ListData)

	http.HandleFunc("/api/list/del", ListDel)

	http.ListenAndServe(":8090", nil)
}

func IndexView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/index.html")
	io.Copy(w, f)
	f.Close()

}

func IndexData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idstr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	mod, err := model.ArticleGet(id)
	if err != nil {
		Fail(w, err.Error())
		return
	}
	buf, _ := json.Marshal(mod)
	w.Write(buf)
}

func ListView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/list.html")
	io.Copy(w, f)
	f.Close()

}

func ListData(w http.ResponseWriter, r *http.Request) {
	mods, err := model.ArticleList()
	if err != nil {
		Fail(w, err.Error())
		return
	}
	buf, _ := json.Marshal(mods)
	w.Write(buf)
}

func ListDel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idstr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	if model.ArticleDel(id) {
		Succ(w, "删除成功")
		return
	} else {
		Fail(w, "删除失败")
		return
	}
}

func Succ(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}

func Fail(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}
