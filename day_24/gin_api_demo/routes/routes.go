/*************
** 定义路由
**************/

package routes

import (
	"go_learn/day_24/gin_api_demo/apps"
	"go_learn/day_24/gin_api_demo/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoute ...
func InitRoute() *gin.Engine {
	route := gin.Default()

	// 加载静态文件
	route.StaticFS("/static", http.Dir("static"))
	// 加载模板
	route.LoadHTMLGlob("templates/*")

	// 首页路由
	route.GET("/", apps.IndexHandler)
	// 404
	route.NoRoute(apps.NoRouteHandler)

	// 定义路由组
	api := route.Group("/api")
	api.Use(middlewares.Auth())
	{
		// 增加用户
		api.POST("/user", apps.AddUserHandler)
		//  获取一个用户
		api.GET("/user/:id", apps.GetOneMemHandler)
		// 获取用户列表
		api.GET("/user", apps.GetMemListHandler)
		// 删除一个用户
		api.DELETE("/user/:id", apps.DeleteMemHandler)
		// 修改一个用户
		api.PUT("/user/:id", apps.EditMemHandler)
	}
	return route
}
