package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	str, _, err := in.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n", string(str))
}
