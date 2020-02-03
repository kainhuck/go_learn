package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker%d start worke%d\n", id, j)
		time.Sleep(3 * time.Second)
		fmt.Printf("Worker%d finish worke%d\n", id, j)
		results <- j * j
	}
}

func main() {
	var (
		jobs    chan int
		results chan int
	)
	jobs = make(chan int, 100)
	results = make(chan int, 100)
	// 三个工人
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	// 9个任务
	for i := 1; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	// for each := range results {
	// 	fmt.Println(each)
	// }
	for i := 1; i < 10; i++ {
		<-results
	}
}
