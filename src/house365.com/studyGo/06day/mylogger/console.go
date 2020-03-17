package mylogger

import (
	"fmt"
	"time"
)

type LogLevel uint16

type ConsoleLogger struct {
	Level LogLevel
}

func NewLog(levelstr string) ConsoleLogger {
	level, err := parseLogLevel(levelstr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}
func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		levelStr := getLogString(lv)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n", now, levelStr, fileName, funcName, lineNo, msg)
	}
}
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(msg string) {
	c.log(INFO, msg)
}

func (c ConsoleLogger) Warning(msg string) {
	c.log(WARNING, msg)
}

func (c ConsoleLogger) Error(msg string) {
	c.log(ERROR, msg)
}

func (c ConsoleLogger) Fatal(msg string) {
	c.log(FATAL, msg)
}
