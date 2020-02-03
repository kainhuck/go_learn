package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {
	url := "https://so.gushiwen.org/authors/"

	// method 1
	pageHTML, err := getHTML(url)
	if err != nil {
		panic(err)
	}
	doc, err := htmlquery.Parse(strings.NewReader(pageHTML))
	if err != nil {
		panic(err)
	}

	// method 2
	// doc, err := htmlquery.LoadURL(url)
	// if err != nil {
	// 	panic(err)
	// }

	// nodes, err := htmlquery.QueryAll(doc, "//div[@class='cont']/a")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(nodes)
	// for _, each := range nodes {
	// 	fmt.Println(each.Data)
	// }

	list := htmlquery.Find(doc, "//div[@class='cont']/a[@href]")
	for _, n := range list {
		fmt.Println(htmlquery.InnerText(n)) // output inner text
	}

	hrefList := htmlquery.Find(doc, "//div[@class='cont']/a/@href")
	for _, n := range hrefList {
		fmt.Println(htmlquery.InnerText(n))
	}

}

// get HTML
func getHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	HTML := string(pageBytes)
	return HTML, nil
}
