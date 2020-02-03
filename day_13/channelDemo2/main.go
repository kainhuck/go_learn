package main

import (
	"sync"
	"time"
)

var (
	ch chan int // 定义全局通道
	wg sync.WaitGroup
)

// 从通道中获取数据
func get() {
	defer wg.Done()
	time.Sleep(time.Second * 4)
	<-ch
}

// 向通道中放数据
func put() {
	defer wg.Done()
	ch <- 1
}

func main() {
	ch = make(chan int)
	wg.Add(2)
	go get()
	go put()
	wg.Wait()
}
