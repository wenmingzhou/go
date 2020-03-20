package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string //日志文件保存的路径
	filename    string //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		filename:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 50000),
	}
	err = f1.initFIle() //按照文件路径和文件名打开文件
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *FileLogger) initFIle() error {
	fullFileName := path.Join(f.filePath, f.filename)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}

	f.fileObj = fileObj
	f.errFileObj = errFileObj
	//开启后台的go去写日志
	for i := 0; i < 1; i++ {
		go f.writeLogBackGround()
	}
	return nil

}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}

	return fileInfo.Size() >= f.maxFileSize

}
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {

	nowStr := time.Now().Format("20060102130405")

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      //拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //拼接一个日志备份
	file.Close()                                           //xx.log.bak201908031709

	os.Rename(logName, newLogName)

	fmt.Println(newLogName)
	//需要切割日志文件
	//1 关闭当前的日志文件
	//2 备份一下 rename
	//3 打开一个新的日志文件
	//4 将打开的新日志文件对象赋值给 f.fileObj

	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new file failed,err:%v\n", err)
		return nil, err
	}
	return fileObj, nil
}

func (f *FileLogger) writeLogBackGround() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s \n", logTmp.timestamp, getLogString(logTmp.level),
				logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				//如果要记录的日志大于等于error级别,要在err日志文件中再记录一遍
				fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s \n", logTmp.timestamp, getLogString(logTmp.level),
					logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)

			}
		default:
			time.Sleep(time.Second)
		}
	}
}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)

		//先把日志发送到通道中
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now,
			line:      lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default: //把日志扔掉保证不出现阻塞
		}

	}
}
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)

}

func (f *FileLogger) Info(msg string) {
	f.log(INFO, msg)

}

func (f *FileLogger) Warning(msg string) {
	f.log(WARNING, msg)

}

func (f *FileLogger) Error(msg string) {
	f.log(ERROR, msg)

}

func (f *FileLogger) Fatal(msg string) {
	f.log(FATAL, msg)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
