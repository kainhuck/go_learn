package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default() // Default 使用 Logger 和 Recovery 中间件
	// engine := gin.New() // 不使用默认中间件
	engine.LoadHTMLGlob("templates/**/*")
	// engine.SecureJsonPrefix(")]}',\n")
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// engine.MaxMultipartMemory = 8 << 20  // 8 MiB
	application(engine)
	engine.Run(":8000")
}
