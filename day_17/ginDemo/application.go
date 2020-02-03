package main

import "github.com/gin-gonic/gin"

func application(engine *gin.Engine) {
	engine.GET("/hello", helloGetHandler)
	engine.POST("/hello", helloPostHandler)
	engine.PUT("/hello", helloPutHandler)
	engine.DELETE("/hello", helloDeleteHandler)
	engine.GET("/querystring", queryStrHandler)
	engine.POST("/postform", postFormHandler)
	engine.GET("/asciijson", asciiJSONHandler)
	engine.GET("/index", indexHandler)
	engine.GET("/posts/index", pindexHandler)
	engine.GET("/users/index", uindexHandler)
	engine.GET("/jsonp", jsonpHandler)
	engine.POST("/login", loginHandler)
	engine.POST("/form_post", formPostHandler)
	engine.GET("/purejson", pureJSONHandler)
	engine.POST("/post", postHandler)
}
