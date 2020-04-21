package taillog

import (
	"fmt"
	"house365.com/studyGo/logagent/etcd"
	"time"
)

var tskMgr *tailLogMgr

//tailTask 管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry // struct c:/tmp/nginx.log   web_log
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, //把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲区通道
	}
	//fmt.Println(tskMgr)
	for _, logEntry := range logEntryConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
	go tskMgr.run()
}

//监听自己的newConfChan,有了新的配置过来之后就做对应的处理
//1.配置新增
//2.配置删除
//3.配置变更

func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println("新的配置来了", newConf)
		default:
			time.Sleep(time.Second)

		}
	}
}

//向外暴露一个函数，tskMgr的newConfChan
func NewConfChan() chan []*etcd.LogEntry {
	return tskMgr.newConfChan
}
