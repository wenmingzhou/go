package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	//3 与客户端通信
	var tmp [128]byte
	for { //同一个人多次聊天
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
func main() {
	//1 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start server on 127.0.0.1:20000 failed", err)
		return
	}
	//2 等待被人来跟我建立连接
	for { //多个连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept failed", err)
			return
		}
		go processConn(conn)
	}

}
