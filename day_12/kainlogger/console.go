package kainlogger

import (
	"fmt"
	"time"
)

// 用于控制台输出日志

// 定义控制台日志结构体
type ConsoleLogger struct {
	Level LogLevel // 开关
}

func NewConsoleLogger(levelStr string) *ConsoleLogger {
	// 根据传入的字符串，将其转化成对应的等级
	level, err := parseLevel(levelStr)
	handleErr(err, "解析等级出错")
	return &ConsoleLogger{
		Level: level,
	}
}

func (c *ConsoleLogger) enable(level LogLevel) bool {
	return c.Level <= level
}

// 日志输出格式
// [时间] [等级] [文件名:函数名:行号] msg

func (c *ConsoleLogger) log(level LogLevel, format string, args ...interface{}) {
	if c.enable(level) {
		msg := fmt.Sprintf(format, args...)
		time := time.Now().Format("2006-02-03 15:04:0500")
		levelStr, err := unparseLevel(level)
		handleErr(err, "反向解析出错")
		funcName, fileName, lineNo := getDetail(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", time, levelStr, fileName, funcName, lineNo, msg)
	}
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DEBUG, format, args...)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	c.log(TRACE, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(INFO, format, args...)
}

func (c *ConsoleLogger) Warning(format string, args ...interface{}) {
	c.log(WARNING, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ERROR, format, args...)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FATAL, format, args...)
}
