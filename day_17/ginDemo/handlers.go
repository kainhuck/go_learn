package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloGetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "Get",
	})
}

func helloPostHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "POST",
	})
}

func helloPutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "Put",
	})
}

func helloDeleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "Delete",
	})
}

func queryStrHandler(c *gin.Context) {
	username := c.DefaultQuery("username", "root")
	password := c.Query("password")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}

func postFormHandler(c *gin.Context) {
	username := c.DefaultPostForm("username", "root")
	password := c.PostForm("password")
	jsonData := make(map[string]interface{})
	jsonData["username"] = username
	jsonData["password"] = password
	c.JSON(http.StatusOK, jsonData)
}

func asciiJSONHandler(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}
	c.AsciiJSON(200, data)
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Main page",
	})
}

func pindexHandler(c *gin.Context) {
	c.HTML(200, "posts/index.tmpl", gin.H{
		"title": "post Main page",
	})
}

func uindexHandler(c *gin.Context) {
	c.HTML(200, "users/index.tmpl", gin.H{
		"title": "user main page",
	})
}

func jsonpHandler(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	c.JSONP(200, data)
}

func loginHandler(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		if form.User == "root" && form.Password == "123" {
			c.JSON(200, gin.H{
				"status": "you are login",
			})
		} else {
			c.JSON(401, gin.H{
				"status": "unauthorizedrror",
			})
		}
	}
}

func formPostHandler(c *gin.Context) {
	username := c.DefaultPostForm("username", "root")
	password := c.PostForm("password")

	c.JSON(200, gin.H{
		"status":   "posted",
		"username": username,
		"password": password,
	})
}

func pureJSONHandler(c *gin.Context) {
	// 通常，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。
	// 如果要按字面对这些字符进行编码，则可以使用 PureJSON。
	// Go 1.6 及更低版本无法使用此功能。
	c.PureJSON(200, gin.H{
		"html": "<b>hello world<b>",
	})
}

func postHandler(c *gin.Context) {
	id := c.Query("id")
	name := c.DefaultQuery("name", "tutu")
	page := c.PostForm("page")
	message := c.PostForm("message")
	fmt.Printf("id: %s, name: %s, page: %s, message: %s\n", id, name, page, message)
}
