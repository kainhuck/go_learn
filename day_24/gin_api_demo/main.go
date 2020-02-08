package main

import (
	"go_learn/day_24/gin_api_demo/db"
	"go_learn/day_24/gin_api_demo/libs"
	"go_learn/day_24/gin_api_demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	defer db.Conns.Close()
	gin.SetMode(libs.Conf.Read("site", "runmode"))
	route := routes.InitRoute()
	route.Run(":" + libs.Conf.Read("site", "httpport"))
}
