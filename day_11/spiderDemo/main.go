package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	qqEmail = `\d+@qq.com`
)

// 写个爬虫吧
func getEmail() {
	// 1. 发送GET请求
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1013319691")
	handleErr(err, "http请求错误")
	defer resp.Body.Close() // 关闭请求

	// 2. 获取页面内容
	pageSource, err := ioutil.ReadAll(resp.Body)
	handleErr(err, "页面内容解析错误")
	// fmt.Println(string(pageSource))

	// 3. 提取数据
	re, err := regexp.Compile(qqEmail)
	handleErr(err, "正则解析错误")
	result := re.FindAllStringSubmatch(string(pageSource), -1)
	// fmt.Println(result)

	// 4. 循环遍历打印
	for _, v := range result {
		fmt.Printf("email: %s\n", v[0])
	}
}

// 处理error
func handleErr(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main() {
	getEmail()
}
