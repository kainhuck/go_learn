/*
** 数字解析就是将字符串转化为数字,因为go语言中字符串到数字不能强制转化
** 需要使用到strconv包
 */

package main

import (
	"fmt"
	"strconv"

	"github.com/kainhuck/fancylog"
)

var log = fancylog.NewConsoleLogger("Debug")

func main() {
	// 使用 ParseFloat，这里的 64 表示解析的数的位数, 这个数字不影响解析出来的类型,都是float64
	f, err := strconv.ParseFloat("1.23", 64)
	handleError(err)
	fmt.Printf("%T -> %v\n", f, f)

	// 在使用 ParseInt 解析整型数时， 例子中的参数 0 表示自动推断字符串所表示的数字的进制。 64 表示返回的整型数是以 64 位存储的。
	i, err := strconv.ParseInt("123", 0, 64)
	handleError(err)
	fmt.Printf("%T -> %v\n", i, i)

	// ParseInt 会自动识别出字符串是十六进制数。
	d, err := strconv.ParseInt("0x1c8", 0, 64)
	handleError(err)
	fmt.Println(d)

	// ParseUint 也是可用的。
	u, err := strconv.ParseUint("789", 0, 64)
	handleError(err)
	fmt.Println(u)
	
	// Atoi 是一个基础的 10 进制整型数转换函数
	k, err := strconv.Atoi("135")
	handleError(err)
	fmt.Println(k)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
