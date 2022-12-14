//对某一个数据库的操作

package dao

import (
	"awesomeProject/model"
	"log"
)

func InsertUser(u model.User) (err error) {
	//当此处以参变量表示了err后，在return时就不必要再return一个变量了。
	_, err = DB.Exec("insert into user (username ,password) values(?,?)", u.Username, u.Password)
	return //单一个return
}
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select username,password from  user  where username =? ", name)
	if err = row.Err(); row.Err() != nil {
		//这里的Err是啥哦
		return
	}
	err = row.Scan(&u.Username, &u.Password)
	return
}

func ChangePass(username string, newpass string) (err error) {
	_, err = DB.Exec("update user set password =? where username = ?", newpass, username)
	if err != nil {
		log.Printf("change error:%v", err)
		return
	}
	return
}
