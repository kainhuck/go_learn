package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reAuthorInfo = `<a href="/authorv_(.+?)\.aspx">(.+?)</a>`
)

func main() {
	url := "https://so.gushiwen.org/authors/"
	pageSource, err := getPageSource(url)
	if err != nil {
		panic(err)
	}

	re, err := regexp.Compile(reAuthorInfo)
	if err != nil {
		fmt.Printf("complie Error,err:%v\n", err)
		return
	}
	result := re.FindAllStringSubmatch(pageSource, -1)
	for _, each := range result {
		link := "https://so.gushiwen.org/authorv_" + each[1] + ".aspx"
		fmt.Printf("link: %s name: %s\n", link, each[2])
	}
	// fmt.Println(pageSource)
}

func getPageSource(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GetError,err:%v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	if code != 200 {
		fmt.Printf("Request failed:statusCode:%d\n", code)
		return "", errors.New("StatusCodeError")
	}
	fmt.Printf("Get success,statusCode:%d\n", code)

	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Get pageSource failed,err:%v\n", err)
		return "", err
	}

	pageStr := string(pageBytes)
	return pageStr, nil
}
