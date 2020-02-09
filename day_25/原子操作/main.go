/**********************************************
** 对一个数据进行原子操作,防止多个线程同时操作一个数据
***********************************************/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64 // 定一个待操作的变量

	var wg sync.WaitGroup // 定义等待组

	// 启动50个线程,每个线程加1000次
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				// ops++ // 加一
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(ops) // 应该是50000
}
