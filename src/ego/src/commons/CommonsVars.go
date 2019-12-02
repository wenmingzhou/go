package commons

import "github.com/gorilla/mux"

var (
	//初始化路由
	Router              = mux.NewRouter()
	CurrentPath         = "http://localhost:8888/"
	HEADER_CONTENT_TYPE = "Content-Type"
	JSON_HEADER         = "application/json;charset=utf-8"
)
