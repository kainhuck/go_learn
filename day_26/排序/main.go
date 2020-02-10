/**********************
** go语言中自带的排序方法
***********************/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. 排序字符串切片
	strs := []string{"c", "r", "w", "a", "u"}
	sort.Strings(strs)
	fmt.Printf("Strings:%v\n", strs)

	// 2. 排序数字切片
	ints := []int{3, 2, 5, 1, 6, 2, 4, 5, 7, 3}
	sort.Ints(ints)
	fmt.Printf("Ints:%v\n", ints)

	// 检查一个切片是否有序
	a := sort.IntsAreSorted(ints)
	fmt.Println(a)
}
