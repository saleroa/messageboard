package api

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

// 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		utils.RespParamError(c)
		return
	}
	u, err := service.SearchUserByUserName(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error :%v", err)
		utils.ResInternalErr(c)
		return
	}

	if u.Username != "" {
		utils.ResNormalErr(c, 300, "用户已存在")
		return
	}

	err = service.Creatruser(model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		utils.ResInternalErr(c)
	}
	utils.RespOk(c)

}

var User string

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		utils.RespParamError(c)
		return
	}
	u, err := service.SearchUserByUserName(username)

	if err != nil {

		if err == sql.ErrNoRows {
			//检测是否是根本不存在这一行数据
			utils.ResNormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			//log就是日志，打印出有时间的消息
			utils.ResInternalErr(c)
			return
		}
		return
	}
	if u.Password != password {
		utils.ResNormalErr(c, 20001, "密码错误")
	}
	User = username
	c.SetCookie("name", username, 0, "", "/", false, true) //每个不同的user的value不一样，各个的cookie都不一样u
	utils.RespOk(c)

}

// 修改密码
// 逻辑是，通过uername确定，在加密码验证，再改密码
func Pass(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		utils.RespParamError(c)
		return
	}
	_, err := service.SearchUserByUserName(username)

	if err != nil {

		if err == sql.ErrNoRows {
			//检测是否是根本不存在这一行数据
			utils.ResNormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			//log就是日志，打印出有时间的消息
			utils.ResInternalErr(c)
			return
		}
		return
	}
	newpass := c.PostForm("newpass")
	err = service.Change(username, newpass)
	if err != nil {
		utils.ResNormalErr(c, 300, "change error")
		return
	}
	utils.RespOk(c)
}
