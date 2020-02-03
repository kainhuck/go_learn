package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	// traversalString()
	// chageString()
	// sqrtDemo()
	// test()
	// switchDemo()
	// gotoDemo()
	// chengfaBiao()
	// getSumOfArray()
	// findEightIndex()
	// sliceDemo()
	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)
	// fmt.Printf("%d\n", len(a))
	// sortDemo()
	// for i := 0; i < 20; i++ {
	// 	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	// 	fmt.Println(key)
	// }
	// mapTest()
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])

}
func traversalString() {
	s := "hello世界"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
	fmt.Printf("%d\n", len(s))
}
func chageString() {
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "红萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '白'
	fmt.Println(string(runeS2))
}
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func test() {
	wordSum := 0
	s := "hello 浙江理工大学"
	for _, r := range s {
		if len(string(r)) == 3 {
			wordSum++
		}
	}
	fmt.Println(wordSum)
}

func switchDemo() {
	s := 75
	switch {
	case s < 60:
		fmt.Println("不及格")
	case s < 80:
		fmt.Println("良好")
	case s < 100:
		fmt.Println("优秀")
	default:
		fmt.Println("ok")
	}
}

func gotoDemo() {
	for {
		for i := 0; i < 100; i++ {
			if i == 19 {
				goto breakTag
			}
		}

	}
breakTag:
	fmt.Println("推出循环")
}

func chengfaBiao() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func getSumOfArray() {
	sum := 0
	array1 := [...]int{1, 3, 5, 7, 8}
	for _, num := range array1 {
		sum += num
	}
	fmt.Println(sum)
}

func findEightIndex() {
	nums := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == 8 {
				fmt.Printf("(%d,%d)\n", nums[i], nums[j])
			}
		}
	}
}

func sliceDemo() {
	// slice1 := []int{1,2,3,4}
	// fmt.Printf("len=%d, cap=%d", len(slice1), cap(slice1))
	//切片再切片
	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))
	b := a[1:3]
	fmt.Printf("b:%v type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	c := b[1:5]
	fmt.Printf("c:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))
	d := make([]int, 3)
	fmt.Printf("%d %d\n", len(d), cap(d))

	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	aa := []string{"成都", "重庆"}
	citySlice = append(citySlice, aa...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]

	var ad = make([]string, 5, 10)
	fmt.Printf("c:%v type:%T len:%d  cap:%d\n", ad, ad, len(ad), cap(ad))
	x := make([]int, 0, 10)
	fmt.Println(x)
}

func sortDemo() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Sort(sort.IntSlice(a[:]))
	fmt.Println(a)
}

func mapTest() {
	/*
		写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
	*/
	testStr := "how do you do"
	strSlice := strings.Split(testStr, " ")
	worldMap := make(map[string]int)
	for _,v := range strSlice{
		times,ok := worldMap[v]
		if !ok{
			times = 0
		}
		times++
		worldMap[v] = times
	}
	fmt.Println(worldMap)
}
