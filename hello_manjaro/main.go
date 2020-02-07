package main

import "fmt"

func main() {

	// a := 1
	// b := 2
	// a, b = b, a // 快速交换两个数
	// fmt.Println(a, b)

	// 类型断言

	// var a interface{}
	// a = 1
	// b, ok := a.(bool)
	// fmt.Printf("b: %v, type of b: %T, a: %v, type of a: %T, ok: %v\n", b, b, a, a, ok)

	var a *int
	fmt.Printf("%#v\n", &a)
}
