package main

import "fmt"

/*使用“面向对象”的思维方式编写一个学生信息管理系统。
学生有id、姓名、年龄、分数等信息
程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能*/

type student struct {
	id    int
	name  string
	age   int8
	score int8
}

var (
	stuList []*student
)

// 构造函数
func newStu(id int, name string, age int8, score int8) *student {
	return &student{
		id, name, age, score,
	}
}

// 展示学生列表
func showStudent() {
	for _, stu := range stuList {
		fmt.Printf("%v\n", stu)
	}
}

// 添加学生
func addStudent(stu *student) {
	stuList = append(stuList, stu)
}

// 编辑学生信息
func (s *student) setInfo(id int, name string, age int8, score int8) {
	s.id = id
	s.age = age
	s.name = name
	s.score = score
}

// 删除学生
func removeStu(stu *student) {
	index := 0
	for _, each := range stuList {
		if each == stu {
			break
		}
		index++
	}
	stuList = append(stuList[:index], stuList[index+1:]...)
}

func main() {
	var tempStu *student
	showStudent()
	fmt.Println("添加学生中")
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("stu%02d", i+1)
		age := i + 20
		score := 60 + i
		stuu := newStu(i+1, name, int8(age), int8(score))
		addStudent(stuu)
		if i == 3 {
			tempStu = stuu
		}
	}
	showStudent()
	removeStu(tempStu)
	fmt.Println("以把４号删除")
	showStudent()
}
