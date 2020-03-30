package main

import "net/http"

func f1(w http.ResponseWriter, r *http.Request) { //w 响应体   r 请求
	str := `<h1 style="color:red">hello 南京</h1>`
	w.Write([]byte(str))
}
func main() {
	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil) //启动9090端口web服务
}
