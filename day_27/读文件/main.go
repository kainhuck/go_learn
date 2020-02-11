package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/kainhuck/fancylog"
)

var log = fancylog.NewConsoleLogger("debug")

// 读文件可能会引发异常,专门定义一个处理异常的函数
func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}

func main() {
	file := "./poeam"
	// 最基本的文件读取任务或许就是将文件内容读取到内存中。
	// 只是简单的一次性都读出来，不做别的任何操作
	// poeam, err := ioutil.ReadFile(file)
	// handleError(err)
	// fmt.Println(string(poeam))

	// *************************************************************
	// **
	// *************************************************************

	// 您通常会希望对文件的读取方式和内容进行更多控制。 对于这个任务，首先使用 Open 打开一个文件，以获取一个 os.File 值
	f, err := os.Open(file)
	handleError(err)

	// 从文件的开始位置读取一些字节。 最多允许读取 5 个字节，但还要注意实际读取了多少个。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	handleError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 你也可以 Seek 到一个文件中已知的位置，并从这个位置开始 Read
	o2, err := f.Seek(6, 0)
	handleError(err)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2)
	handleError(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// 例如，io 包提供了一个更健壮的实现 ReadAtLeast，用于读取上面那种文件
	o3, err := f.Seek(6, 0)
	handleError(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	handleError(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内建的倒带，但是 Seek(0, 0) 实现了这一功能
	_, err = f.Seek(0, 0)
	handleError(err)

	// ***********************************************************************
	// **
	// ***********************************************************************

	// bufio 包实现了带缓冲的读取， 这不仅对于很多小的读取操作能够提升性能， 也提供了很多附加的读取函数。 bufio 包实现了带缓冲的 Reader， 这对于提高读取大量小文件的效率， 以及它提供的其他读取方法可能都有用
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	handleError(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
