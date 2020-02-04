package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() // Default 使用 Logger 和 Recovery 中间件
	// engine := gin.New() // 不使用默认中间件
	engine.LoadHTMLGlob("templates/**/*")
	// engine.SecureJsonPrefix(")]}',\n")
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// engine.MaxMultipartMemory = 8 << 20  // 8 MiB

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 记录到文件
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 自定义路由日志的格式
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	application(engine)
	// engine.Run(":8000")
	// 自定义http配置
	s := &http.Server{
		Addr:           ":8000",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
