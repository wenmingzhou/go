package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1 与server 端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial 127.0.0.1:20000 failed", err)
		return
	}
	//2 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请说话:")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

	conn.Close()

}
