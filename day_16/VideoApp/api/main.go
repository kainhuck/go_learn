package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Application 路由
func Application() *httprouter.Router {
	router := httprouter.New() // 构造函数

	router.POST("/user", CreateUserHandler)
	router.POST("/user/:username", LoginHandler)

	return router
}

func main() {
	app := Application()
	http.ListenAndServe(":8000", app)
}
