package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile1() {
	fileObj, err := os.Open("./test.txt")

	if err != nil {
		fmt.Printf("打开文件出错，%s\n", err)
		return
	}

	defer fileObj.Close()

	var tmp [128]byte
	var content []byte

	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("读取文件出错，%s", err)
			break
		}
		// fmt.Printf("已读取%d字节数据\n", n)
		// fmt.Println(string(tmp[:n]))
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

func readFromFile2() {
	fileObj, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("打开文件出错，%s", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				fmt.Println(line)
			} else {
				fmt.Println("结束")
				break
			}
		}
		if err != nil {
			fmt.Printf("文件读取出错，%s", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFile3() {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read the file failed, err: ", err)
		return
	}
	fmt.Println(string(content))
}

func writeFile1() {
	fileObj, err := os.OpenFile("test.py", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("打开文件出错，%s", err)
		return
	}
	defer fileObj.Close()
	str := "print('hello world')"
	fileObj.WriteString(str)
	fmt.Println("文件写入成功")
}

func writeFile2() {
	fileObj, err := os.OpenFile("test2.py", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("打开文件出错，%s", err)
		return
	}
	writer := bufio.NewWriter(fileObj)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("print('hello %d')\n", i)
		writer.WriteString(str)
	}
	writer.Flush()
	fmt.Println("文件写入成功")
}

func writeFile3() {
	err := ioutil.WriteFile("test3.py", []byte("print('hello xianxian')"), 0666)
	if err != nil {
		fmt.Printf("文件写入出错，%s", err)
		return
	}
}

func main() {
	// readFromFile1()
	// readFromFile2()
	// readFromFile3()
	// writeFile1()
	// writeFile2()
	writeFile3()
}
