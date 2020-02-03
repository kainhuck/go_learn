package main

import (
	"fmt"
	"sync"
)

// channel测试

var (
	can chan int
	wg  sync.WaitGroup
)

func noBufChannel() {
	can = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-can
		fmt.Println("接收到值", x)
	}()
	can <- 10
	fmt.Println("向通道中发送一个10")
}

func bufChannel() {
	can = make(chan int, 1)
	can <- 10
	fmt.Println("向通道中发送一个10")
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-can
		fmt.Println("接收到值", x)
	}()
	can <- 20
	fmt.Println("向通道中发送一个20")
	close(can)
}

func main() {
	bufChannel()
	wg.Wait()
}
