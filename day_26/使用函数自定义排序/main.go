/**************************
** 使用函数自定义排序
** 本例子按照字符串长度进行排序
** 我们为该类型实现了
** sort.Interface 接口的
** Len、Less 和 Swap 方法，
** 这样我们就可以使用 sort 包
** 的通用 Sort 方法了，
***************************/

package main

import (
	"fmt"
	"sort"
)

// SortByLength ...
type SortByLength []string

func (a SortByLength) Len() int           { return len(a) }
func (a SortByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func main() {
	str := []string{"tutu", "xsy", "kangkang"}
	sort.Sort(SortByLength(str))
	fmt.Println(str)
}
