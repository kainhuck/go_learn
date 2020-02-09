package main

import (
	"errors"
	"fmt"
)

type myError struct {
	value interface{}
	prob  string
}

func main() {
	// _, err := f1(0)
	// fmt.Println(err.Error())
	// _, err := f2(0)
	// fmt.Println(err.Error())
	// fmt.Printf("%T, %T\n", f1, f2)
	fmt.Printf("%T, %T\n", errors.New("dasd"), &myError{})
}

// 实现 Error() 方法的结构体的指针就是 error 类型
func (e *myError) Error() string {
	return fmt.Sprintf("Your value is %v, error is %v\n", e.value, e.prob)
}

func f1(arg int) (int, error) {
	if arg == 0 {
		return 0, errors.New("别给我0") //errors.New 使用给定的错误信息构造一个基本的 error 值
	}
	return 1, nil
}

func f2(arg int) (int, error) {
	if arg == 0 {
		return 0, &myError{value: 0, prob: "别给我0"}
	}
	return 1, nil
}
