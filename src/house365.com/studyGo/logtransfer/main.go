package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"house365.com/studyGo/logtransfer/conf"
	"house365.com/studyGo/logtransfer/es"
	"house365.com/studyGo/logtransfer/kafka"
)

//将日志数据从kafka取出来发往ES
func main() {
	// 0 加载配置文件
	cfg := new(conf.LogTransferCfg)
	err := ini.MapTo(&cfg, "./conf/cfg.ini") //要传指针
	if err != nil {
		fmt.Printf("ini config,err:%v\n", err)
		return
	}
	fmt.Println(cfg)

	fmt.Println("init kafka success.")
	// 1  初始化ES
	// 1.1初始化一个ES连接的client
	// 1.2对外提供一个往ES写入数据的函数
	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Printf("init es failed,err:%v\n", err)
		return
	}
	//2 初始化kafak连接
	//2.1 连接kafka，创建分区的消费者
	//2.2 每个分区的消费者分别取出数据 通过sendToEs()将数据发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}

	select {}
}
