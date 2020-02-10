/*****************************
** 组合函数
**  strings 切片的组合函数示例
******************************/

package main

import (
	"fmt"
	"strings"
)

// Index 返回目标字符串 t 在 vs 中第一次出现位置的索引， 或者在没有匹配值时返回 -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include 如果目标字符串 t 存在于切片 vs 中，则返回 true
func Include(vs []string, t string) bool {
	for _, v := range vs {
		if v == t {
			return true
		}
	}
	return false
}

// Any 如果切片 vs 中的任意一个字符串满足条件 f，则返回 true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All 如果切片 vs 中的所有字符串都满足条件 f，则返回 true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter 返回一个新的切片，新切片由原切片 vs 中满足条件 f 的字符串构成
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map 返回一个新的切片，新切片的字符串由原切片 vs 中的字符串经过函数 f 映射后得到
func Map(vs []string, f func(string) string) []string {
	vsf := make([]string, len(vs))
	for i, v := range vs {
		vsf[i] = f(v)
	}
	return vsf
}

func main() {
	strs := []string{"tutu", "huhu", "yaoyao", "kangkang", "love", "tutuandhuhu"}
	fmt.Println(Index(strs, "tutu"))
	fmt.Println(Include(strs, "kainhuck"))
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "t")
	}))
	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "t")
	}))
	fmt.Println(Filter(strs, func(s string) bool {
		return strings.HasPrefix(s, "t")
	}))
	fmt.Println(Map(strs, func(s string) string {
		return s+"_love_forever~"
	}))
}
