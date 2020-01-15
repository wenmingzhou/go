package control

import (
	"io"
	"net/http"
	"os"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/index.html")
	io.Copy(w, f)
	f.Close()

}

func ListView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/list.html")
	io.Copy(w, f)
	f.Close()

}

func ArticleAddView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/add.html")
	io.Copy(w, f)
	f.Close()

}
