package main

import "fmt"

func main() {
	var m1 map[int]interface{}
	m1 = make(map[int]interface{}, 1)
	m1[1] = 1
	fmt.Println(len(m1))

}
