/*********
** 互斥锁
**********/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var ops uint64 // 定一个待操作的变量

	var mutex = &sync.Mutex{}

	var wg sync.WaitGroup // 定义等待组

	// 启动50个线程,每个线程加1000次
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				ops++ // 加一
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(ops) // 应该是50000
}
