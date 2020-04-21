package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"house365.com/studyGo/logagent/conf"
	"house365.com/studyGo/logagent/etcd"
	"house365.com/studyGo/logagent/kafka"
	"house365.com/studyGo/logagent/taillog"
	"sync"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

func main() {
	//0 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}

	//1 初始化kafak连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	//2 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")
	//2.1  从etcd中获取日志收集项的配置信息

	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n", err)
		return
	}
	fmt.Println(logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v \n", index, value)
	}

	//2.2  派一个哨兵去监视日志收件项的变化(有变化及时通知我的logAgent实现热加载)

	//3 收集日志发往kafka
	//3.1 循环每一个收集项,创建TailObj
	var wg sync.WaitGroup
	taillog.Init(logEntryConf)
	NewConfChan := taillog.NewConfChan() //从taillog包中获取对外暴露的通道
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, NewConfChan) //哨兵发现最新的配置信息会配置上面的通道
	wg.Wait()
}
