package mylogger

import (
	"fmt"
	"time"
)

// 在终端写日志相关内容

// ConsoLogger 日志结构体
type ConsoLogger struct {
	Level LogLevel
}

// NewConsoLoger 构造函数
func NewConsoLoger(levelStr string) ConsoLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoLogger{
		Level: level,
	}
}

// 控制级别
func (c ConsoLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

// log 记录日志方法  传入格式化参数和空接口(任意个和任意类型参数都可) 仿照的print()传参
func (c ConsoLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		// 拼接参数
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funceName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funceName, fileName, lineNo, msg)
	}
}

// Debug 方法
func (c ConsoLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

// Info 方法
func (c ConsoLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

// Warning 方法
func (c ConsoLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

// Error 方法
func (c ConsoLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

// Fatal 方法
func (c ConsoLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
