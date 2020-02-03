package main

import "fmt"

type person struct {
	name string
	age  uint16
}

func main() {
	var p person
	p.name = "tutu"
	p.age = 17
	fmt.Printf("%s\n", p)
}

func (p person) String() string {
	return fmt.Sprintf("name: %s, age: %d\n", p.name, p.age)
}
