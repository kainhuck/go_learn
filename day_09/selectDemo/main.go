package main

import "fmt"

// selectDemo

func main() {
	var (
		can chan int
	)
	can = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case can <- 1:
			fmt.Println("ok")
		case x := <-can:
			fmt.Println(x)
		}
	}
}
