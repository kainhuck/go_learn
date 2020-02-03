package kainlogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 向文件中写入

// 定义一个结构体，将所有操作都定义为该结构体的方法
type FileLogger struct {
	Level       LogLevel
	path        string
	fileName    string
	errFileObj  *os.File
	fileObj     *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, path, fileName string, maxFileSize int64) *FileLogger {
	level, err := parseLevel(levelStr)
	handleErr(err, "解析出错")
	f := &FileLogger{
		Level:       level,
		path:        path,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	f.initFile()
	return f
}

// 打开日志文件
func (f *FileLogger) initFile() {
	// 普通文件全名
	fullFileName := path.Join(f.path, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err, "日志文件打开失败")
	f.fileObj = fileObj

	// 错误日志文件
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err, "日志文件打开失败")
	f.errFileObj = errFileObj
}

func (f *FileLogger) enable(level LogLevel) bool {
	return f.Level <= level
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	handleErr(err, "读取文件状态出错")
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) *os.File {
	fileInfo, err := file.Stat()
	handleErr(err, "获取文件状态错误")

	oldName := path.Join(f.path, fileInfo.Name())
	// 1. 关闭当前日志文件
	file.Close()
	// 2. 将文件重命名 xxx.log -> xxx.log.bak12434213
	newName := fmt.Sprintf("%s.bak%d", oldName, time.Now().Unix())
	os.Rename(oldName, newName)
	// 3. 创建一个新的日志文件
	fileObj, err := os.OpenFile(oldName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err, "日志文件打开失败")

	// 4. 将新创建日志文件赋值给f.fileObj
	return fileObj
}

// 日志输出格式
// [时间] [等级] [文件名:函数名:行号] msg

func (f *FileLogger) log(level LogLevel, format string, args ...interface{}) {
	if f.enable(level) {

		// 检查文件是否需要切割
		if f.checkSize(f.fileObj) {
			f.fileObj = f.splitFile(f.fileObj)
		}

		msg := fmt.Sprintf(format, args...)
		time := time.Now().Format("2006-02-03 15:04:0500")
		levelStr, err := unparseLevel(level)
		handleErr(err, "反向解析出错")
		funcName, fileName, lineNo := getDetail(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", time, levelStr, fileName, funcName, lineNo, msg)
		if level >= ERROR {
			// 检查文件是否需要切割
			if f.checkSize(f.errFileObj) {
				f.errFileObj = f.splitFile(f.errFileObj)
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", time, levelStr, fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DEBUG, format, args...)
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	f.log(TRACE, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(INFO, format, args...)
}

func (f *FileLogger) Warning(format string, args ...interface{}) {
	f.log(WARNING, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ERROR, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FATAL, format, args...)
}
