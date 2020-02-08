/******************
** 逻辑业务Handler
*******************/

package apps

import (
	"go_learn/day_10/mylogger"
	"go_learn/day_24/gin_api_demo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var log = mylogger.NewConsoleLogger("Debug")

// AddUserHandler ...
func AddUserHandler(c *gin.Context) {
	// 初始化一个对象
	m := new(models.Member)
	m.Username = c.Request.FormValue("username")
	m.Password = c.Request.FormValue("password")

	if id, err := m.AddMember(); err != nil {
		// 失败
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err.Error())
	} else {
		m.ID = uint(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    m,
		})

	}
}

// GetOneMemHandler ...
func GetOneMemHandler(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))
	mem, err := models.OneMember(mid)
	if err != nil {
		// 失败
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    mem,
		})
	}
}

// GetMemListHandler ...
func GetMemListHandler(c *gin.Context) {
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))

	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	list, count, err := models.ListMember(page, pageSize, filters...)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "SUCCESS",
			"data":      list,
			"count":     count,
			"page_size": pageSize,
			"current":   page,
		})
	}

}

// DeleteMemHandler ...
func DeleteMemHandler(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))

	if n, err := models.DeleteMember(mid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}

// EditMemHandler ...
func EditMemHandler(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))
	m := new(models.Member)
	m.Username = c.Request.FormValue("username")
	m.Password = c.Request.FormValue("password")
	if n, err := m.UpdateMember(mid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}
