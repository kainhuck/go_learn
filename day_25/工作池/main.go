/**********
** 工作池
***********/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, result chan<- int) {
	defer wg.Done()
	for j := range jobs { // 持续从工作池里取任务
		fmt.Printf("woker_%d start job\n", id)
		time.Sleep(500 * time.Millisecond)
		result <- j * 2
		fmt.Printf("woker_%d finished job\n", id)
	}
}

func main() {
	const jobNum = 10                // 定义10个任务数量
	jobs := make(chan int, jobNum)   // 定义job通道
	result := make(chan int, jobNum) // 定义result通道

	// 开启3个工人
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, result)
	}

	// 往任务通道里塞任务
	for i := 0; i < jobNum; i++ {
		jobs <- i
	}
	close(jobs)

	// 取数据
	for a := 0; a < jobNum; a++ {
		j := <-result
		fmt.Printf("result: %d\n", j)
	}

	wg.Wait()
}
