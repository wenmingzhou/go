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

func Fail(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}
