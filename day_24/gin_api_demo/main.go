package main

import (
	"bytes"
	"fmt"
	"go_learn/day_24/gin_api_demo/libs"
)

func main() {
	// conf := libs.Conf
	// fmt.Println(conf.Read("mysql", "password"))
	secret := libs.Conf.Read("api", "apisecrect")
	key := libs.Conf.Read("api", "apikey")

	b := bytes.Buffer{}
	b.WriteString("app_key=")
	b.WriteString(key)
	b.WriteString("&app_secret=")
	b.WriteString(secret)
	b.WriteString("&method=POST")
	// b.WriteString(method)
	b.WriteString("&ts=1508304822")
	// b.WriteString(ts)
	fmt.Println(libs.Md5([]byte(b.String())))
}
