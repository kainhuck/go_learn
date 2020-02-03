package main

import (
	"fmt"
	"sync"
)

/*
创建两个channel cha1 cha2
1. 向cha1中放100个数
2. 将cha1中的100个数取出平方后加入cha2
3. 将cha2中的数打印出来
*/

func main() {
	var (
		cha1 chan int
		cha2 chan int
		wg   sync.WaitGroup
	)
	cha1 = make(chan int)
	cha2 = make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			cha1 <- i
		}
		close(cha1)
	}()
	go func() {
		defer wg.Done()
		for {
			i, ok := <-cha1
			if !ok {
				break
			}
			cha2 <- i * i
		}
		close(cha2)
	}()

	for x := range cha2 {
		fmt.Println(x)
	}
	wg.Wait()
}
