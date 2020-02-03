package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var temp []byte
	var content []byte
	temp = make([]byte, 1024)
	for {
		n, err := file.Read(temp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("出错了", err)
			break
		}
		// fmt.Print(string(temp))
		content = append(content, temp[:n]...)
	}
	fmt.Println(string(content))
}
