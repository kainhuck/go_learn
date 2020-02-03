package main

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

示例1：

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例2：

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例3：

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	i := 0
	max := 0
	charList := []rune(s)

	for index, c := range charList {
		for n := i; n < index; n++ {
			if charList[n] == c {
				i = n + 1
			}
		}
		if index-i+1 > max {
			max = index - i + 1
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("sadasddsefad"))
}
