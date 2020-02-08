/************
** 首页Handler
*************/

package apps

import (
	"go_learn/day_24/gin_api_demo/libs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler ...
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "default.html", gin.H{
		"title": libs.Conf.Read("site", "appname"),
	})
}
