package main

import "github.com/gin-gonic/gin"

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


type formA struct{
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct{
	Bar string `json:"bar" xml:"bar" binding:"required"`
}