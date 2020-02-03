package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())   // 年
	fmt.Println(now.Month())  // 月
	fmt.Println(now.Day())    // 日
	fmt.Println(now.Hour())   // 时
	fmt.Println(now.Minute()) // 分
	fmt.Println(now.Second()) // 秒

	// 获取时间戳，本质是将Time类型转化成时间戳
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 时间戳，精确到纳秒
	fmt.Println(timestamp1)
	fmt.Println(timestamp2)

	// 定时器
	// ticker := time.Tick(time.Second)
	// for i := range ticker {
	// 	fmt.Println(i)
	// }

	// 时间格式化
	fmt.Println(now.Format("2006-01-02 03:04:05 PM Mon Jan"))

	// 解析字符串格式的时间
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	meet, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-01 20:00:00", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(meet)

	//时间操作
	// 1. 加
	latter := now.Add(time.Hour) // 当前时间加一个小时
	fmt.Println(latter)
	// 2. 减
	long := now.Sub(meet)
	fmt.Println(long)
}
