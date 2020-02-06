package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type userInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:12345678@(localhost)/img?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&userInfo{})

	// // 创建记录
	// u1 := userInfo{1, "huhu", "boy", "tutu"}
	// u2 := userInfo{2, "tutu", "girl", "huhu"}
	// db.Create(&u1)
	// db.Create(&u2)

	// 查询记录
	u := new(userInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
	uu := new(userInfo)
	db.Find(uu, "hobby=?", "huhu")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "yaoyao")
	
	// 删除
	db.Delete(&u)
}
