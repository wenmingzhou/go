package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"house365.com/studyGo/logagent/kafka"
)

//专门从日志文件中收集日志的模块

//TailTask 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init() //根据路径去打开对应的日志
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在就保存
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config) //打开文件
	if err != nil {
		fmt.Println("tail file failed,err:", err)
	}
	go t.run() //直接去采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines: //从tailobj的通道中一行一行的读取日志数据
			//3.2 发往Kafka
			//kafka.SendToKafka(t.topic,line.Text)  //函数调用函数
			//先把日志数据发送到一个通道中
			kafka.SendToChan(t.topic, line.Text)
			// kafka 包中有单独的goroutine去取日志数据发到kafka
		}
	}
}
