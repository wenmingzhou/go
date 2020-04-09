package main

import (
	"fmt"
	"house365.com/studyGo/lagagent/kafka"
	"house365.com/studyGo/lagagent/taillog"
	"math/rand"
	"os"
	"strings"
	"time"
)

func run() {
	// 1 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2 发送到kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		time.Sleep(20)
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func writelog() {
	for {
		fd, _ := os.OpenFile("my.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		fd_content := strings.Join([]string{"====== ", "=====", GetRandomString(6), "\n"}, "")
		buf := []byte(fd_content)
		fd.Write(buf)
		fd.Close()

	}
}
func main() {
	//1 初始化kafak连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")
	//2 打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success.")

	go writelog()
	run()
}
