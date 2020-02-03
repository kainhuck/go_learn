package main

import (
	"fmt"
	"os"
)

// 学生管理系统结构体版本

// 定义一个全局管理员
var studentMgr stuMgr

// 菜单
func showMenu() {
	fmt.Println(`
	～～欢迎使用学生管理系统～～
	1. 展示所有信息
	2. 增加学生
	3. 删除学生
	4. 编辑学生
	5. 退出
	`)
}

func main() {
	// 实例化管理员对象
	studentMgr = stuMgr{
		allStu: make(map[int64]*student, 40),
	}

	// 死循环
	for {
		// 1. 打印菜单
		showMenu()

		// 2. 用户选择
		var choice int
		fmt.Print("请输入选择>> ")
		fmt.Scanln(&choice)

		// 3. 根据用户选择执行对应方法
		switch choice {
		case 1:
			studentMgr.showStudents()
		case 2:
			studentMgr.addStudent()
		case 3:
			studentMgr.deleteStudent()
		case 4:
			studentMgr.editStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Printf("输入有误！！")
		}

	}
}
