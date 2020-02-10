package main

import (
	"go_learn/day_10/mylogger"
)

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

	// fmt.Printf("%shello world\n", "\033[30m")
	// fmt.Printf("%shello world\n", "\033[31m")
	// fmt.Printf("%shello world\n", "\033[32m")
	// fmt.Printf("%shello world\n", "\033[33m")
	// fmt.Printf("%shello world\n", "\033[34m")
	// fmt.Printf("%shello world\n", "\033[35m")
	// fmt.Printf("%shello world\n", "\033[36m")
	// fmt.Printf("%shello world\n", "\033[37m")

	// fmt.Printf("%shello world\n", "\033[40m")
	// fmt.Printf("%shello world\n", "\033[41m")
	// fmt.Printf("%shello world\n", "\033[42m")
	// fmt.Printf("%shello world\n", "\033[43m")
	// fmt.Printf("%shello world\n", "\033[44m")
	// fmt.Printf("%shello world\n", "\033[45m")
	// fmt.Printf("%shello world\n", "\033[46m")
	// fmt.Printf("%shello world\n", "\033[47m")

	// fmt.Printf("%shello world\n", "\033[0m")
	log := mylogger.NewConsoleLogger("Debug")
	log.Debug("debug")
	log.Trace("trace")
	log.Info("info")
	log.Warning("warning")
	log.Error("error")
	log.Fatal("fatal_%s", "test  ")

	var a []int
	a = append(a, 1)
	log.Debug("%v", a)
}
