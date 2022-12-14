//dao.go是链接数据库，然后把链接的信息封装装起来
//封装的DB

package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //dao层必备
	"log"                              //这个是啥
)

var DB *sql.DB //？？？？

func Initdb() {
	db, err := sql.Open("mysql", "root:301044825@tcp(127.0.0.1:3306)/message_board")
	if err != nil {
		log.Printf("connect error :%v", err) //？？
	}
	DB = db //？？？？
	fmt.Println(db.Ping())
}
