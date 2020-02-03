package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")
	application(engine)
	engine.Run(":8000")
}
