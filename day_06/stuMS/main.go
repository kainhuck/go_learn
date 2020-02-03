package main

import (
	"fmt"
	"os"
)

// 学生管理系统的实现
// 1. 展示所有学生	2. 增加学生	3. 删除学生

// 定义学生结构体
type student struct {
	id   int    // 学号
	name string // 姓名
}

func newStu(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

// 定义全局变量存储班级信息
var (
	allStu map[int]*student
)

func showAllStu() {
	for id, student := range allStu {
		fmt.Printf("学号：%d，姓名: %s\n", id, student.name)
	}
}

func addStu() {
	var (
		id   int
		name string
	)
	// 1. 用户输入学号姓名
	fmt.Print("请输入学号>> ")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名>> ")
	fmt.Scanln(&name)
	// 2. 构造学生结构体
	newStu := newStu(id, name)
	// 3. 加入到班级map
	allStu[id] = newStu
}

func deleteStu() {
	var (
		id int
	)
	// 1. 输入要删除的学生的id
	fmt.Print("请输入学号>> ")
	fmt.Scanln(&id)
	// 2. 删除对应学生id
	delete(allStu, id)
}

func main() {
	allStu = make(map[int]*student, 40)
	for {
		// 1. 打印菜单
		fmt.Println(`
	～～～欢迎使用学生管理系统～～～
	1. 展示所有学生
	2. 增加学生
	3. 删除学生
	4. 退出
	`)

		// 2. 用户输入
		var choice int
		fmt.Print("请输入你的选择>> ")
		fmt.Scanln(&choice)

		// 3. 执行对应函数
		switch choice {
		case 1:
			showAllStu()
		case 2:
			addStu()
		case 3:
			deleteStu()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入出错")
		}
	}
}
