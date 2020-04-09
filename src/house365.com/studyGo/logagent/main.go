package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"house365.com/studyGo/logagent/conf"
	"house365.com/studyGo/logagent/kafka"
	"house365.com/studyGo/logagent/taillog"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	// 1 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2 发送到kafka
			kafka.SendToKafka(cfg.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

func main() {
	//0 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}

	/*cfg,err :=ini.Load("./conf/config.ini")
	fmt.Println("Data Path:", cfg.Section("kafka").Key("address").String())
	*/

	//1 初始化kafak连接
	err = kafka.Init([]string{cfg.Address})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")
	//2 打开日志文件准备收集日志
	err = taillog.Init(cfg.FileName)
	if err != nil {
		fmt.Printf("init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success.")

	run()
}
