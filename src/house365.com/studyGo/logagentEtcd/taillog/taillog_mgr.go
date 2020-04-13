package taillog

import (
	"house365.com/studyGo/logagent/etcd"
)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	//tskMap map[string]*TailTask
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry: logEntryConf, //把当前的日志收集项配置信息保存起来
	}
	//fmt.Println(tskMgr)
	for _, logEntry := range logEntryConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
