package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在就保存
		Poll:      true,
	}
	tails, err := tail.TailFile(filename, config) //打开文件

	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for true {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
