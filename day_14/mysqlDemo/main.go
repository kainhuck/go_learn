package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB // 全局数据库对象
)

// 定义数据库ORM
type user struct {
	id   int
	name string
	age  int
}

// 初始化连接数据库
func initDB() (err error) {
	// 链接DSN: data source name
	dsn := "root:12345678@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库,这里只是验证参数不是真正连接
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("参数验证错误,err:%v\n", err)
		return
	}

	// 尝试连接
	err = db.Ping()
	if err != nil {
		fmt.Printf("连接失败,err:%v\n", err)
		return
	}
	return
}

// 查询单条记录
func query() {
	sql := "select id,name,age from user where id=?"

	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sql, 2).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("查询失败,err:%v\n", err)
		return
	}
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", u.id, u.name, u.age)
}

// 查询多条记录
func queryMultiRow() {
	sql := "select * from user where id>?"

	rows, err := db.Query(sql, 0)
	if err != nil {
		fmt.Printf("查询出错,err%v\n", err)
		return
	}
	// 一定要关闭
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("取得结果出错,err:%v\n", err)
			return
		}
		fmt.Printf("Id: %d, Name: %s, Age: %d\n", u.id, u.name, u.age)
	}

}

// 插入数据
func insert() {
	sql := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sql, "yaoyao", 17)
	if err != nil {
		fmt.Printf("insert err: %v\n", err)
		return
	}

	// 新插入记录的ID
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("Get id err: %v\n", err)
		return
	}
	fmt.Printf("插入成功 id:%d\n", id)
}

// 更新记录
func update() {
	sql := "update user set age=? where id=?"

	ret, err := db.Exec(sql, 20, 3)
	if err != nil {
		fmt.Printf("update err:%v\n", err)
	}

	// 受影响的行数
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rows err:%v\n", err)
	}
	fmt.Printf("受影响的行数是:%d\n", rows)
}

// 删除记录
func delete() {
	sql := "delete from user where id=?"
	ret, err := db.Exec(sql, 3)
	if err != nil {
		fmt.Printf("delete err:%v\n", err)
	}

	// 受影响的行数
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rows err:%v\n", err)
	}
	fmt.Printf("受影响的行数是:%d\n", rows)
}

func main() {
	err := initDB()
	if err != nil {
		return
	}
	// fmt.Println("数据库连接成功")
	// query()
	queryMultiRow()
	// insert()
	// update()
	// delete()
}
