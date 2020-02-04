package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// LoginForm ...
type LoginForm struct {
	User     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//文件上传
var upfilePath = "./upfile"

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

// Person ...
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// Root ...
type Root struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// 绑定html复选框
type myForm struct {
	Colors []string `form:"colors[]"`
}

// 绑定uri
type user struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

// 绑定查询字符串或表单数据
type queryData struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}
