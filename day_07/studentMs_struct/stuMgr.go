package main

import "fmt"

// 学生结构体
type student struct {
	id   int64
	name string
}

// 管理员结构体
type stuMgr struct {
	allStu map[int64]*student
}

// 管理员方法
// 1. 展示所有学生
func (s stuMgr) showStudents() {
	// 判断学生是否为空
	if len(s.allStu) == 0 {
		fmt.Println("暂时还没有学生！")
		return
	}

	for _, stu := range s.allStu {
		fmt.Printf("学号：%d\t, 姓名：%s\n", stu.id, stu.name)
	}
}

// 2. 增加学生
func (s stuMgr) addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号>> ")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名>> ")
	fmt.Scanln(&name)

	s.allStu[id] = &student{
		id:   id,
		name: name,
	}
	fmt.Println("添加成功！")
}

// 3. 删除学生
func (s stuMgr) deleteStudent() {
	var (
		id int64
	)
	fmt.Print("请输入学号>> ")
	fmt.Scanln(&id)
	delete(s.allStu, id)
	fmt.Println("删除成功！")
}

// 4. 编辑学生
func (s stuMgr) editStudent() {
	var (
		id int64
	)
	fmt.Print("请输入学号>> ")
	fmt.Scanln(&id)
	stu, ok := s.allStu[id]
	if ok {
		// 修改
		fmt.Printf("你要修改的学生学号是：%d, 姓名是: %s\n", stu.id, stu.name)
		var (
			newID   int64
			newName string
		)
		fmt.Print("请输入新学号>> ")
		fmt.Scanln(&newID)
		fmt.Print("请输入新姓名>> ")
		fmt.Scanln(&newName)
		stu.id = newID
		stu.name = newName

		fmt.Println("修改学生成功！")
	} else {
		fmt.Println("没有这个学生！")
	}
}
