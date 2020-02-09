/*************************************************
** 定时器 是当你想要在未来某一刻执行一次时使用的 -
** 打点器 则是为你想要以固定的时间间隔重复执行而准备的。
**************************************************/

package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 创建打点器
	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool) // 创建一个通道用于关闭本程序

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Ticker at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
