/*************
** 字符串格式化
**************/
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)           // {1 2}
	fmt.Printf("%+v\n", p)          // {x:1 y:2} %+v 的格式化输出内容将包括结构体的字段名
	fmt.Printf("%#v\n", p)          // main.point{x:1, y:2} %#v 根据 Go 语法输出值，即会产生该值的源码片段。
	fmt.Printf("%T\n", p)           // main.point 需要打印值的类型，使用 %T
	fmt.Printf("%t\n", false)       // false 格式化bool值
	fmt.Printf("%d\n", 123)         // 123 格式化整型数有多种方式，使用 %d 进行标准的十进制格式化
	fmt.Printf("%b\n", 14)          // 1110 二进制
	fmt.Printf("%c\n", 33)          // ! 输出给定整数的对应字符。
	fmt.Printf("%x\n", 456)         // 1c8 16进制小写
	fmt.Printf("%X\n", 456)         // 1C8 16进制大写
	fmt.Printf("%f\n", 13.14)       // 13.140000 使用 %f 进行最基本的十进制格式化。
	fmt.Printf("%e\n", 123400000.0) // %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学记数法表示形式
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")           // "string" 使用 %s 进行基本的字符串输出
	fmt.Printf("%q\n", "\"string\"")           // "\"string\""  像 Go 源代码中那样带有双引号的输出，使用 %q
	fmt.Printf("%p\n", &p)                     // 0xc000014120 输出指针的值(地址)
	fmt.Printf("|%6d|%6d|\n", 12, 345)         // |    12|   345| 格式化数字时，您经常会希望控制输出结果的宽度和精度。 要指定整数的宽度，请在动词 “%” 之后使用数字。 默认情况下，结果会右对齐并用空格填充。
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)   // |  1.20|  3.45| 你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  | 要左对齐，使用 - 标志。
	fmt.Printf("|%6s|%6s|\n", "foo", "b")      // |   foo|     b| 你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。 这是基本的宽度右对齐方法。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")    // |foo   |b     | 要左对齐，使用 - 标志。

	s := fmt.Sprintf("a %s", "string")         // 到目前为止，我们已经看过 Printf 了， 它通过 os.Stdout 输出格式化的字符串。 Sprintf 则格式化并返回一个字符串而没有任何输出
	fmt.Printf(s)                              // a string
	fmt.Fprintf(os.Stderr, "an %s\n", "error") // an error 你可以使用 Fprintf 来格式化并输出到 io.Writers 而不是 os.Stdout。
}
