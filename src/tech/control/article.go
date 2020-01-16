package control

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"tech/model"
	"time"
)

type Reply struct {
	Code int         `json:"code"` //200 成功 300 失败 310 输入有误 320 输出有误
	Msg  string      `json:"msg"`  //用户提示
	Data interface{} `json:"data"` //返回数据
}

type Ext struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`
}

func ApiArticleAdd(w http.ResponseWriter, r *http.Request) {

	mod := &model.Article{}
	//err := json.NewDecoder(r.Body).Decode(mod)
	buf, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(buf, mod)
	if err != nil {
		Fail(w, "数据有误"+err.Error())
		return
	}

	mod.Utime = time.Now()
	err = model.ArticleAdd(mod)
	if err != nil {
		Fail(w, "添加失败"+err.Error())
		return
	}
	Succ(w, "添加成功")
	return
}

func ApiArticleEdit(w http.ResponseWriter, r *http.Request) {

	mod := &model.Article{}
	//err := json.NewDecoder(r.Body).Decode(mod)
	buf, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(buf, mod)
	if err != nil {
		Fail(w, "数据有误"+err.Error())
		return
	}

	err = model.ArticleEdit(mod)
	if err != nil {
		Fail(w, "编辑失败"+err.Error())
		return
	}
	Succ(w, "编辑成功")
	return
}

/*func ApiArticleAdd(w http.ResponseWriter, r *http.Request) {
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
}*/

func ListData(w http.ResponseWriter, r *http.Request) {
	mods, err := model.ArticleList()
	if err != nil {
		Fail(w, err.Error())
		return
	}
	Succ(w, "列表", mods)
}

func ApiListPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pi, _ := strconv.Atoi(r.Form.Get("pi"))
	ps, _ := strconv.Atoi(r.Form.Get("ps"))

	count := model.ArticlePageCount()
	if count < 1 {
		Fail(w, "未查询数据", "count < 1")
	}
	mods, _ := model.ArticlePage(pi, ps)
	if len(mods) < 1 {
		Fail(w, "未查询数据", "len(mods) < 1")
	}
	ext := Ext{
		Count: count,
		Items: mods,
	}
	Succ(w, "分页数据", ext)
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

func Succ(w http.ResponseWriter, msg string, data ...interface{}) {
	mod := Reply{
		Code: 200,
		Msg:  msg,
	}
	if len(data) > 0 {
		mod.Data = data[0]
	}
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func Fail(w http.ResponseWriter, msg string, data ...interface{}) {
	mod := Reply{
		Code: 300,
		Msg:  msg,
	}
	if len(data) > 0 {
		mod.Data = data[0]
	}
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
