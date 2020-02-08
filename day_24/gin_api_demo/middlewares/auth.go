/*************
** 权限检查
**************/

package middlewares

import (
	"bytes"
	"go_learn/day_24/gin_api_demo/libs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth 权限验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.FormValue("app_key")
		sign := c.Request.FormValue("sign")
		ts := c.Request.FormValue("ts")
		method := c.Request.Method
		now := time.Now().Unix()

		if key != libs.Conf.Read("api", "apikey") {
			noAuth(c, "Key Error")
			return
		}

		// 时差超过两秒没有权限
		timeCheck, _ := strconv.Atoi(ts)
		if (now - int64(timeCheck)) > 100000000000 {
			noAuth(c, "time out")
			return
		}

		// 判断签名
		if ! Sign(key, ts, method, sign) {
			noAuth(c, "Unauthorized")
			return
		}
		c.Next()
		return
	}
}

// noAuth 验证不通过
func noAuth(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
	c.Abort()
}

// Sign 签名
func Sign(key, ts, method, sign string) bool {
	secret := libs.Conf.Read("api", "apisecrect")

	b := bytes.Buffer{}
	b.WriteString("app_key=")
	b.WriteString(key)
	b.WriteString("&app_secret=")
	b.WriteString(secret)
	b.WriteString("&method=")
	b.WriteString(method)
	b.WriteString("&ts=")
	b.WriteString(ts)
	if libs.Md5([]byte(b.String())) == sign {
		return true
	}
	return false
}
