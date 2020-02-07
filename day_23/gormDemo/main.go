package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定一个结构体对应数据表
type user struct {
	ID     uint // gorm 会默认将ID设为主键
	Name   string
	Age    uint8
	Gender string
	Hobby  string
}

var (
	db *gorm.DB
)

func main() {
	// 1. 链接数据库
	if err := connectDB(); err != nil {
		fmt.Println("Error")
		return
	}

	defer closeDB() // 记得要关闭数据库连接

	// 2. 创建数据表
	// createDB()

	// 3. 插入数据
	// insertDB()

	// 4. 查询数据
	// selectDB()

	// 5. 更新数据
	// updateDB()

	// 6. 删除数据
	// deleteDB()
}

func connectDB() (err error) {
	db, err = gorm.Open("mysql", "root:12345678@(localhost)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	return nil
}

func closeDB() {
	db.Close()
}

func createDB() {
	db.AutoMigrate(&user{})
}

func insertDB() {
	u1 := user{
		ID:     0,
		Name:   "huhu5",
		Age:    17,
		Gender: "male",
		Hobby:  "tutu5",
	}
	ok := db.NewRecord(u1) // 主键为空返回 true
	fmt.Println(ok)
	db.Create(&u1)        // 插入
	ok = db.NewRecord(u1) // 创建`u1`后返回`false`
	// fmt.Println(ok)
	if !ok {
		fmt.Println("插入成功!")
	}
}

func selectDB() {
	var u user
	// 根据主键查询第一条记录
	// db.First(&u)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 随机获得一条数据 ??
	// db.Take(&u)
	//// SELECT * FROM users LIMIT 1;

	// 根据主键查询最后一条记录
	// db.Last(&u)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	// fmt.Printf("%#v\n", u)

	// var users []user
	// // 查询所有的记录
	// db.Find(&users)
	//// SELECT * FROM users;
	// fmt.Println(users)

	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&u, 3)
	//// SELECT * FROM users WHERE id = 3;
	fmt.Printf("%#v\n", u)
}

func updateDB() {
	var u user
	db.First(&u, 3)
	db.Model(&u).Update("name", "huhu3") // 未指定 u 全更新
}

func deleteDB() {
	var u user
	db.Find(&u)
	db.Delete(&u)
}
