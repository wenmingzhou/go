package control

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func IndexView(w http.ResponseWriter,r * http.Request){
	f,err :=os.Open("./src/news/views/index.html")
	if err!=nil{
		fmt.Println("os.Open err",err)
	}
	io.Copy(w,f)
	f.Close()
}

func ListView(w http.ResponseWriter,r * http.Request){
	f,err :=os.Open("./src/news/views/list.html")
	if err!=nil{
		fmt.Println("os.Open err",err)
	}
	io.Copy(w,f)
	f.Close()
}


func ViewArticleAdd(w http.ResponseWriter,r * http.Request){
	f,err :=os.Open("./src/news/views/add.html")
	if err!=nil{
		fmt.Println("os.Open err",err)
	}
	io.Copy(w,f)
	f.Close()
}
