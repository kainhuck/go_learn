package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello() {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	// fmt.Println("hello world")
	// fmt.Println(rand.Int())
	fmt.Printf("hello %d\n", rand.Intn(10))
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		hello()
	}
	wg.Wait()
}
