package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// putOnlyDemo1()
	// getOnlyDemo1()

	ch := make(chan int, 2)
	ch <- 1
	// wg.Add(2)
	putOnlyDemo2(ch)
	getOnlyDemo2(ch)
	getOnlyDemo2(ch)
	getOnlyDemo2(ch)
	// wg.Wait()
}

func putOnlyDemo1() {
	var ch chan<- int //只能存
	ch = make(chan int, 1)
	ch <- 1
	close(ch) // 可以关闭
	// <-ch	// 编译不通过
}

func putOnlyDemo2(ch chan<- int) {
	// defer wg.Done()
	ch <- 2
	// <-ch	// 编译不通过
	close(ch)
}

func getOnlyDemo1() {
	var ch <-chan int // 只能取
	ch = make(chan int, 1)
	// ch <- 1 // 编译不通过,不能向只能取的通道中放值
	// close(ch) // 不能关闭只能取的通道
	<-ch // 会报错,因为通道为空且未关闭

	// 所以这种情况没有意义
}
func getOnlyDemo2(ch <-chan int) {
	// defer wg.Done()
	n, ok := <-ch
	fmt.Println(n, ok)
	// ch <- 2 // 编译不通过,不能向只能取的通道中放值
	// close(ch) // 不能关闭只能取的通道
}
