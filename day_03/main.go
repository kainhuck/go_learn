package main

import "fmt"

type animal interface {
	move()
	roll()
}

type dog struct {
	name string
}

type cat struct {
	name string
}

func (d dog) move() {
	fmt.Printf("%s可以像狗一样跑\n", d.name)
}

func (c cat) move() {
	fmt.Printf("%s可以像猫一样跑\n", c.name)
}

func (d dog) roll() {
	fmt.Println("汪汪")
}

func (c cat) roll() {
	fmt.Println("喵喵")
}

func main() {
	d := dog{
		name: "旺财",
	}
	c := cat{
		name: "小花",
	}
	var x animal
	// d.move()
	// d.roll()
	// c.move()
	// c.roll()
	x = d
	x.move()
	x.roll()
	x = c
	x.move()
	x.roll()
}
