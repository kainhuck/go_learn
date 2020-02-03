package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// douban top250 spider

var (
	reMovieName = `<span class="title">(.+?)</span>`
)

func handleError(err error, reason string) {
	if err != nil {
		fmt.Printf("%v, %s\n", err, reason)
	}
}

// 解析网页
func parseHTML(pageSource string) {
	re, err := regexp.Compile(reMovieName)
	handleError(err, "regexp.Compile")
	result := re.FindAllStringSubmatch(pageSource, -1)
	// fmt.Println(result)
	for _, v := range result {
		fmt.Println(v[1])
	}
}

// 获取网页源码
func getPageSorce(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	handleError(err, "http.NewRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	resp, err := client.Do(req)
	handleError(err, "client.Do")
	pageBytes, err := ioutil.ReadAll(resp.Body)
	handleError(err, "resp.Body")
	return string(pageBytes)
}

func main() {
	url := "http://movie.douban.com/top250"
	pageSource := getPageSorce(url)
	// fmt.Println(pageSource)
	parseHTML(pageSource)
}
