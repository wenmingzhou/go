package commons

//客户端服务器端数据交换模板
type EgoResult struct {
	Status int  //200成功
	Data interface{}
	Msg string
}