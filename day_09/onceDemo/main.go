package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func hello() {
	fmt.Println("hello")
}

func main() {
	for i := 0; i < 3; i++ {
		once.Do(hello)
	}
}
