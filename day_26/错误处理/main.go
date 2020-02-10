/********************************
** 利用panic和recover实现错误的处理
*********************************/

package main

import (
	"errors"

	"github.com/kainhuck/fancylog"
)

var log = fancylog.NewConsoleLogger("Debug")

func main() {
	f1()
}

func f1() {
	defer func() {
		// recover只能在defer中调用
		e := recover() // 捕获panic抛出的异常
		if e, ok := e.(error); ok {
			log.Info("handle this error successful")
		} else {
			log.Error("can not handle this error, type%T\n", e)
		}
	}()

	err := errors.New("this is error")
	panic(err) // 抛出异常
}
