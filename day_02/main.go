package main

import "fmt"

func main() {
	// fmt.Println(intSum(1, 2, 3, 4, 7))
	// fmt.Println(calc(1, 2))
	// fmt.Println(calcV2(1, 2))
	// var f = adder()
	// fmt.Println(f(10)) // 10
	// fmt.Println(f(20)) // 30
	// fmt.Println(f(30)) // 60

	// f1 := adder()
	// fmt.Println(f1(40)) // 40
	// fmt.Println(f1(50)) // 90
	// deferDemo()
	deferDemo2()
}

func intSum(x ...int) int {
	sum := 0

	// for _, v := range x {
	// 	sum += v
	// }

	// fmt.Printf("%T\n",x)	// int类型的切片

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum
}

func calc(x int, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func calcV2(x int, y int) (a int, b int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func deferDemo() {
	fmt.Println("Start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func deferDemo2() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
