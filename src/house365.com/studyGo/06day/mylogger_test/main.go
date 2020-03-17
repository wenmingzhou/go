package main

import (
	"house365.com/studyGo/06day/mylogger"
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
	//log :=mylogger.NewLog("debug")
	log := mylogger.NewFileLogger("debug", "./", "zhoulin.log", 10*1024*1024)
	for {
		id := 10010
		name := "理想"
		log.Debug("这是一条Debug日志,id:%d,name:%s", id, name)
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
