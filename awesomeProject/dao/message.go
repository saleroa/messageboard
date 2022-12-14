package dao

import (
	"awesomeProject/model"
	"log"
)

//从数据库里取出相应的信息

func InsertMessage(m model.Message) (err error) {
	_, err = DB.Exec("insert into message (username ,message) values (?,?)", m.Username, m.Message)
	return
}

var Word model.Message

func GetMessage(username string) (err error) {
	row := DB.QueryRow("select * from message where username = ?", username)
	err = row.Scan("&Word.Username ,&Word.Message ")
	return
}

func ChangeMessage(username string, message string) (err error) {
	_, err = DB.Exec("  update message set message = ? where username = ?  ", message, username)
	if err != nil {
		log.Printf("change error:%v", err)
		return
	}
	return
}
func DeleteMessage(username string) (err error) {
	_, err = DB.Exec("delete from message where usernme =?", username)
	if err != nil {
		log.Printf("delete error :%v", err)
		return
	}
	return
}
