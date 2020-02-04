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
	engine.GET("secureJSON", secureJSONHandler)
	engine.GET("/someJSON", someJSONHandler)
	engine.GET("/moreJSON", moreJSONHandler)
	engine.GET("/someXML", someXMLHandler)
	engine.GET("/someYAML", someYAMLHandler)
	// engine.GET("/someProtoBuf", someProtoBufHandler)
	engine.POST("/uploadSingle", uploadSingleHandler)
	engine.POST("/uploadMulti", uploadMultiHandler)

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", secretsHandler)

	engine.GET("/bindQuery", bindQueryHandler)
	engine.GET("/long_async", longAsyncHandler)
	engine.GET("/long_sync", longSyncHandler)
	engine.GET("multiBind", multiBindHandler)
	engine.GET("/mapargs", mapArgsHandler)
	
}