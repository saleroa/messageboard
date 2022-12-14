package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)
//拿dao层的函数
func AddMessage (m model.Message)(err error){
	err = dao.InsertMessage(m)
	return
}


//注意此处的username要从cookie来
func GainMessage(username string)(err error){
	err = dao.GetMessage(username)
	return
}

func FixMessage(username string ,message string)(err error){
	err = dao.ChangeMessage(username, message)
     return
}

func SubMessage (username string )(err error){
	err = dao.DeleteMessage(username)
	return
}