package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// fileName := "fileDemo.txt"
	// 一次性写入并关闭文件，简单粗暴
	// d1 := []byte("hello\ngolang\n")
	// err:=ioutil.WriteFile(fileName, d1, 0644)
	// check(err)

	// 对于更细粒度的写入，先打开一个文件
	f, err := os.Create("fileDemo2.txt") // 创建一个文件并打开
	check(err)
	defer f.Close() // 记得关闭

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // 调用 Sync 将缓冲区的数据写入硬盘

	// *******************************************************
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()	// 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer

}
