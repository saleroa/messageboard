//api的router是创建的路由组
//router就是对所有的路由组进行组合，
//api层除了router.go以外的包都是不同的功能块

package api

import (
	"github.com/gin-gonic/gin"
)

func Initrouter() {
	r := gin.Default()

	u := r.Group("/user")
	{
		u.POST("/register", Register) //ok
		u.POST("/login", Login)       //ok
		u.PUT("/password", Pass)      //OK
	}

	//message功能的架构

	//不同电脑认证的cookie应该是不同的，
	//然后设置不同的mid，再根据不同的mid取不同的message

	//逻辑还没开始

	m := r.Group("/message")
	{
		//获得自己的message，需要通过model里的mid实现
		m.GET("/message", GetMessage)
		//发送
		m.POST("/message", SendMessage)
		//更新
		m.PUT("/message", PutMessage)
		//删除
		m.DELETE("/message", DelMessage)
	}

	_ = r.Run()
}

//把auth放在每个函数内部的前面。
