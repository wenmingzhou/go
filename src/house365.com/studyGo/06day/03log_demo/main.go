package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

/**需求
 * 支持往不同的地方输出日志
 * 日志分级别
 *  	1.debug
  		2.info
 		3.warning
        4.error
        5.fatal
    日志要支持开关控制
	日志要有时间，行号，文件名，日志级别，日志信息
    日志文件要切割


*/
func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
	}
	log.SetOutput(fileObj)
	for {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 3)
	}
}
