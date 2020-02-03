package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func do(name string, ch chan int, ju chan int) {
	defer wg.Done()
	for {
		_, ok := <-ch
		if !ok {
			fmt.Printf("%s 赢了\n", name)
			return
		}
		flag := rand.Intn(100) // 接球标志
		if flag%13 == 0 {      // 丢球
			fmt.Printf("%s 没接到\n", name)
			close(ch)
			return
		}
		fmt.Printf("%s 接到球\n", name)
		ch <- 1
		ju <- 1
		time.Sleep(time.Second)
	}
}

func main() {
	// 定义并初始化一个无缓冲通道
	ch := make(chan int)
	// close(ch)
	// v, ok := <-ch
	// fmt.Println(v, ok)
	// v, ok = <-ch
	// fmt.Println(v, ok)

	/*
		对关闭通道的理解:
		1. 向一个未关闭的空通道取值会报错,向一个已经关闭的空通道取值不会报错,会返回对应类型的零值
		2. 向一个关闭的非空通道取值不会有异常,可以顺利取到对应的值
		3. 向一个关闭的通道存值会出现异常
		4. 向一个已经满了的通道存值会出现异常(如果没有gorutine取值,否则阻塞),

		总结:
		通道关闭是对 取操作 友好的,一个关闭了的通道是只能取不能存的,而且取完不会报错,而是返回false
	*/

	// 尝试向通道中放值, 报错
	// ch <- 1

	// 尝试向通道中取值, 报错
	// <-ch

	// 尝试关闭通道, 不报错
	// close(ch)

	ju := make(chan int, 1000)

	wg.Add(2)
	go do("Kangkang", ch, ju)
	go do("TuTu", ch, ju)
	ch <- 1 // 发球
	wg.Wait()
	fmt.Printf("此次比赛一共打了%d回合\n", int((len(ju)+1)/2))
}
