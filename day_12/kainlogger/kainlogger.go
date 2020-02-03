package kainlogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 重新写一边自己的log库

// 1. 定义日志的六个等级

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLevel(levelStr string) (LogLevel, error) {
	// 先将levelStr统一转化成大写
	levelStr = strings.ToUpper(levelStr)
	switch levelStr {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("输入的日志等级有误")
	}
}

func unparseLevel(level LogLevel) (string, error) {
	switch level {
	case DEBUG:
		return "DEBUG", nil
	case TRACE:
		return "TRACE", nil
	case INFO:
		return "INFO", nil
	case WARNING:
		return "WARNING", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	default:
		return "UNKNOWN", errors.New("内部出错，这句话一般是不可能输出的")
	}
}

func handleErr(err error, why string) {
	if err != nil {
		fmt.Printf("%s ,err: %v\n", why, err)
	}
}

func getDetail(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)

	if !ok {
		fmt.Println("getDetail failed!")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}

