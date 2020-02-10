/****************
** go语言正则表达式
*****************/

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面我们是直接使用字符串，但是对于一些其他的正则任务， 你需要通过 Compile 得到一个优化过的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 检查是否匹配
	fmt.Println(r.MatchString("peach"))

	// 查找匹配的字符串, 返回地一个匹配的字符串
	fmt.Println(r.FindString("peach punch dsadas"))

	// 这个也是查找首次匹配的字符串， 但是它的返回值是，匹配开始和结束位置的索引，而不是匹配的内容
	fmt.Println(r.FindStringIndex("peach punch, dasdas"))

	// Submatch 返回完全匹配和局部匹配的字符串。 例如，这里会返回 p([a-z]+)ch 和 ([a-z]+) 的信息。
	// 返回第一个匹配的字符串
	fmt.Println(r.FindStringSubmatch("peach punch sads"))
	// 类似的，这个会返回首次完全匹配和局部匹配位置的索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch sads"))

	// 带 All 的这些函数返回全部的匹配项， 而不仅仅是首次匹配项。例如查找匹配表达式全部的项
	// All 同样可以对应到上面的所有函数。
	// 这些函数接收一个正整数作为第二个参数，来限制匹配次数。 负数代表全部
	fmt.Println(r.FindAllString("peach punch dsasd", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach pinch punch sadas", 2))

	// 上面的例子中，我们使用了字符串作为参数， 并使用了 MatchString 之类的方法。 我们也可以将 String 从函数命中去掉，并提供 []byte 的参数。
	fmt.Println(r.Match([]byte("peach")))

	// 创建正则表达式常量时，可以使用 Compile 的变体 MustCompile。 普通的 Compile 不适用于常量，因为它有 2 个返回值。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp 包也可以用来将子字符串替换为为其它值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func 变体允许您使用给定的函数来转换匹配的文本。
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
