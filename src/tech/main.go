package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", IndexView)
	http.ListenAndServe(":8090", nil)
}

func IndexView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./views/index.html")
	io.Copy(w, f)
	f.Close()

}
