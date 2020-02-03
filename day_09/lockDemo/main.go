package main

import (
	"fmt"
	"sync"
)

// lockDemo

var (
	lock sync.Mutex     // 定义全局锁
	x    = 0            // 全局变量
	wg   sync.WaitGroup // 等待组
)

// 操作函数
func f1() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		lock.Lock() // 对临界资源上锁
		x++
		lock.Unlock() // 对临界资源解锁
	}
}

func main() {
	// 开启三个goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go f1()
	}
	wg.Wait()
	fmt.Println(x)
}
