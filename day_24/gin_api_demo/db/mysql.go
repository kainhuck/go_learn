/******************************
** 连接 mysql
** 提供一个可操作数据库的对象 Conns
** 导入该文件时就已经初始化完成
*******************************/

package db

import (
	"database/sql"
	"go_learn/day_10/mylogger"
	"go_learn/day_24/gin_api_demo/libs"

	// ...
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Conns ...
	Conns *sql.DB
	log   = mylogger.NewConsoleLogger("Debug")
)

func init() {
	var err error
	// 读取配置
	host := libs.Conf.Read("mysql", "host")
	port := libs.Conf.Read("mysql", "port")
	username := libs.Conf.Read("mysql", "username")
	password := libs.Conf.Read("mysql", "password")
	database := libs.Conf.Read("mysql", "database")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	Conns, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = Conns.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)
}
