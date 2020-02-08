package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

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

func secureJSONHandler(c *gin.Context) {
	names := []string{"tutu", "kain", "huhu"}
	c.SecureJSON(200, names) // while(1);["tutu","kain","huhu"]
}

func someJSONHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello JSON",
		"status":  "success",
	})
}

func moreJSONHandler(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "tutu"
	msg.Message = "hey"
	msg.Number = 2020
	c.JSON(http.StatusOK, msg)
}

func someXMLHandler(c *gin.Context) {
	c.XML(200, gin.H{
		"message": "hey XML",
		"status":  http.StatusOK,
	})
}

func someYAMLHandler(c *gin.Context) {
	c.YAML(200, gin.H{
		"message": "hey YAML",
		"status":  http.StatusOK,
	})
}

// func someProtoBufHandler(c *gin.Context) {
// 	reps := []int64{int64(1), int64(2)}
// 	label := "test"
// 	// protobuf 的具体定义写在 testdata/protoexample 文件中。
// 	data := &protoexample.Test{
// 		Label: &label,
// 		Reps:  reps,
// 	}
// 	// 请注意，数据在响应中变为二进制数据
// 	// 将输出被 protoexample.Test protobuf 序列化了的数据
// 	c.ProtoBuf(http.StatusOK, data)
// }

func uploadSingleHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// 上传文件至指定目录
	dst := path.Join(upfilePath, file.Filename)
	c.SaveUploadedFile(file, dst)

	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 	curl -X POST http://localhost:8080/upload \
	//   -F "file=@/Users/appleboy/test.zip" \
	//   -H "Content-Type: multipart/form-data"
}

func uploadMultiHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		// 上传文件至指定目录
		dst := path.Join(upfilePath, file.Filename)
		c.SaveUploadedFile(file, dst)
	}
	c.String(200, fmt.Sprintf("%d files uploaded!", len(files)))
}

func secretsHandler(c *gin.Context) {
	// 获取用户，它是由 BasicAuth 中间件设置的
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

func bindQueryHandler(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("========Only Bind By Query String=========")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "ok")
}

func longAsyncHandler(c *gin.Context) {
	// 当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。

	// 创建一个context副本
	cCp := c.Copy()

	go func() {
		// 模拟一个长任务
		time.Sleep(5 * time.Second)

		// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func longSyncHandler(c *gin.Context) {
	// 模拟一个长任务
	time.Sleep(5 * time.Second)

	// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
	log.Println("Done! in path " + c.Request.URL.Path)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func multiBindHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	if err := c.ShouldBind(&objA); err == nil {
		c.String(200, "bind in formA")
	} else if err := c.ShouldBind(&objB); err == nil {
		c.String(200, "bind in formB")
	}
}

func mapArgsHandler(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func bindJSONHandler(c *gin.Context) {
	var user Root

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if user.Username != "tutu" || user.Password != "huhu" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged in",
	})
}

func bindXMLHandler(c *gin.Context) {
	var user Root

	if err := c.ShouldBindXML(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if user.Username != "tutu" || user.Password != "huhu" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "loggin successfully!",
	})
}

func bindFormHandler(c *gin.Context) {
	var form Root

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if form.Username != "tutu" || form.Password != "huhu" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "loggin successfully",
	})
}

func bindCheckBoxHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{
		"color": fakeForm.Colors,
	})
}

func bindURIHandler(c *gin.Context) {
	var u user

	if err := c.ShouldBindUri(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": u.Name,
		"id":   u.ID,
	})
}

func bindQueryStrHandler(c *gin.Context) {
	var qd queryData
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	if c.ShouldBind(&qd) == nil {
		log.Println(qd.Name)
		log.Println(qd.Address)
		log.Println(qd.Birthday)
	}

	c.String(200, "success")
}

func writeCookieHandler(c *gin.Context) {

	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", http.SameSiteLaxMode, false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "cookie writed",
	})
}

func readCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"cookie": "NONE",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cookie": cookie,
	})
}

func clearCookieHandler(c *gin.Context) {
	c.SetCookie("gin_cookie", "test", -1, "/", "localhost", http.SameSiteLaxMode, false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "cookie cleared",
	})
}

func anyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func noRouteHandler(c *gin.Context) {
	fmt.Println("Not Found 404")
	c.HTML(http.StatusNotFound, "404.html", nil)
}

func showHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "show something",
	})
}

func listHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "list something",
	})
}

func middlewareDemoHandler(c *gin.Context) {
	name := c.MustGet("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func redirectHTTPHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.sogo.com/")
}

func testHandler(c *gin.Context) {
	name := c.Request.PostFormValue("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
