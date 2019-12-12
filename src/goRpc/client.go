package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	var req float32 //请求值
	req = 3
	var resp *float32 //返回值
	//同步调用的方式
	err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)

	/*var respSync * float32
	synCall :=client.Go("MathUtil.CalculateCircleArea",req,&respSync,nil)
	replayDone :=<-synCall.Done
	fmt.Println(replayDone)
	fmt.Println(*respSync)*/
}
