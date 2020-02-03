package main

import (
	"fmt"
	"go_learn/requests"
)

func main() {
	url := "https://www.gushiwen.org/app/weixin.png"
	headers := map[string]string{
		`User-Agent`: `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36`,
	}
	r, err := requests.Get(url, headers)
	if err != nil {
		panic(err)
	}
	// file, err := os.OpenFile("1.png", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprint(file, r.Content)
	fmt.Printf("状态码: %d\n", r.StatusCode)

	// doc, err := htmlquery.Parse(strings.NewReader(r.Text))
	// if err != nil {
	// 	panic(err)
	// }
	// titleList := htmlquery.Find(doc, `//div[@class="main3"]/div[@class="left"]/div[@class="sons"]/div[@class="cont"]/p[1]//b`)
	// for index, each := range titleList {
	// 	fmt.Printf("%d %s\n", index+1, htmlquery.InnerText(each))
	// }
}
