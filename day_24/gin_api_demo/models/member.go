/********************
** 数据模型
*********************/

package models

import (
	"go_learn/day_24/gin_api_demo/db"
	"log"
	"strconv"
)

// Member ...
type Member struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// AddMember ...
func (m *Member) AddMember() (id int64, err error) {
	// 插入到数据库
	res, err := db.Conns.Exec("INSERT INTO demo_users(username, password) VALUES (?, ?);", m.Username, m.Password)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

// DeleteMember ...
func DeleteMember(id int) (n int64, err error) {
	res, err := db.Conns.Exec("DELETE FROM demo_users WHERE id=?;", id)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	n, err = res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

// UpdateMember ...
func (m *Member) UpdateMember(id int) (n int64, err error) {
	res, err := db.Conns.Exec("UPDATE demo_users SET username=?, password=? WHERE id=?", m.Username, m.Password, id)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	n, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}

// ListMember ...
func ListMember(page, pageSize int, filters ...interface{}) (list []Member, count int64, err error) {
	list = make([]Member, 0)
	where := "WHERE 1=1"
	if len(filters) > 0 { // 有约束,构造where子句
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	// 构造limit子句
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)

	// 执行SQL
	rows, err := db.Conns.Query("SELECT id, username, password FROM demo_users " + where + " LIMIT " + limit + ";")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer rows.Close()

	count = 0
	// 构造数据对象,加入到list中
	for rows.Next() {
		var tempMem Member
		rows.Scan(&tempMem.ID, &tempMem.Username, &tempMem.Password)
		list = append(list, tempMem)
		count++
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
		return
	}
	return

}

// OneMember ...
func OneMember(id int) (m Member, err error) {
	m.ID = 0
	m.Username = ""
	m.Password = ""
	err = db.Conns.QueryRow("SELECT id, username, password FROM demo_users WHERE id=? LIMIT 1;", id).Scan(&m.ID, &m.Username, &m.Password)
	return
}
