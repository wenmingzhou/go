package control

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tech/model"
	"time"
)

func ApiArticleAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mod := &model.Article{}
	mod.Title = r.Form.Get("title")
	mod.Author = r.Form.Get("author")
	mod.Content = r.Form.Get("content")
	mod.Hits, _ = strconv.Atoi(r.Form.Get("hits"))
	mod.Utime = time.Now()
	err := model.ArticleAdd(mod)
	if err != nil {
		Fail(w, "添加失败"+err.Error())
		return
	}
	Succ(w, "添加成功")
	return
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
