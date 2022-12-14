package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}

func Creatruser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

func Change(username string, newpass string) error {
	err := dao.ChangePass(username, newpass)
	return err
}
