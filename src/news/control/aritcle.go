package control

import (
	"encoding/json"
	"net/http"
	"news/model"
	"strconv"
	"time"
)



//单条
func IndexData(w http.ResponseWriter,r * http.Request)  {
	r.ParseForm()
	idstr :=r.Form.Get("id")
	id,_ :=strconv.ParseInt(idstr,10,64)
	mod,err :=model.ArticleGet(id)
	if err!=nil{
		Fail(w,err.Error())
		return
	}
	buf,_:=json.Marshal(mod)
	w.Write(buf)
}

//多条
func ListData(w http.ResponseWriter,r * http.Request)  {
	r.ParseForm()
	mod,err :=model.ArticleList()
	if err!=nil{
		Fail(w,err.Error())
		return
	}
	Succ(w,"列表",mod)

}

//删除
func ListDel(w http.ResponseWriter,r * http.Request)  {
	r.ParseForm()
	idstr :=r.Form.Get("id")
	id,_ :=strconv.ParseInt(idstr,10,64)
	if model.ArticleDel(id) {
		Succ(w,"删除成功")
		return
	}

	Fail(w, "删除失败")
	return

}

func ApiArticleAdd(w http.ResponseWriter,r * http.Request) {
	r.ParseForm()
	mod 		:=&model.Article{}
	mod.Title 	=r.Form.Get("title")
	mod.Author 	=r.Form.Get("author")
	mod.Content =r.Form.Get("content")
	mod.Hits    =r.Form.Get("hits")
	mod.Utime 	=time.Now()
	err :=model.ArticleAdd(mod)
	if err ==nil{
		Succ(w,"添加成功")
		return
	}
	Fail(w,"添加失败"+err.Error())
	return
}



func Fail(w http.ResponseWriter,msg string,data ...interface{})  {
	mod:=Replay{
		Code:300,
		Msg:msg,
	}
	if len(data)>0{
		mod.Data =data[0]
	}
	buf,_ :=json.Marshal(mod)
	w.Write(buf)
}

func Succ(w http.ResponseWriter,msg string,data ...interface{})  {
	mod:=Replay{
		Code:200,
		Msg:msg,
	}
	if len(data)>0{
		mod.Data =data[0]
	}
	buf,_ :=json.Marshal(mod)
	w.Write(buf)
}

type Replay struct {
	Code int  		`json:"code"` //200成功  300 失败  310输入有误 320 输出有误
	Msg  string 	 `json:"msg"`
	Data interface{} `json:"data"`
}
