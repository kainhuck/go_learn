package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwlock sync.RWMutex   // 定义全局读写锁
	lock   sync.Mutex     // 定义全局普通锁
	wg     sync.WaitGroup // 定义等待组
	x      = 0            // 定义临界资源
)

// 写函数
func write() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		// lock.Lock()	// 4.47113
		rwlock.Lock() // 3.30194
		x++
		// lock.Unlock()
		rwlock.Unlock()
	}
}

// 读函数
func read() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		// lock.Lock()
		rwlock.RLock()
		fmt.Println(x)
		// lock.Unlock()
		rwlock.RUnlock()
	}
}

func main() {
	start := time.Now()
	// 开启写goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write()
	}

	// 开启读线程
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
