package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func first(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte(`hello world`))
}

func index(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	//name :=r.Form.Get("name")
	name :=r.Form["name"]

	fmt.Println(name)
	f,err:=os.Open("./src/view/index.html")
	if err!=nil{
		fmt.Println("os.open err ",err)
	}
	io.Copy(w,f)
}

func indexjs(w http.ResponseWriter,r *http.Request)  {
	f,err:=os.Open("./src/js/1.js")
	if err!=nil{
		fmt.Println("os.open err ",err)
	}
	io.Copy(w,f)
}
func main() {


	http.HandleFunc(`/`,index)
	http.HandleFunc(`/1.js`,indexjs)
	http.HandleFunc(`/first`,first)
	http.ListenAndServe(":8080",nil)
}
