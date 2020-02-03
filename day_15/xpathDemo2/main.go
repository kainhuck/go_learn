package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func poemSpider() {
	url := "https://so.gushiwen.org/authors/"

	// 1. requests
	reps, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer reps.Body.Close()
	pageBytes, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		panic(err)
	}
	pageHTML := string(pageBytes)

	// 2. parse
	doc, err := htmlquery.Parse(strings.NewReader(pageHTML))
	if err != nil {
		panic(err)
	}

	link := htmlquery.Find(doc, "//div[@class='cont']/a/@href")

	// 3. print
	for _, each := range link {
		fmt.Println(htmlquery.InnerText(each))
	}
}

func main() {
	poemSpider()
}
