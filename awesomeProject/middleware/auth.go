package middleware

import (
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
)

// 这里需要读取cookie

func Auth(c *gin.Context) bool {
	_, err := c.Cookie("name")
	if err != nil {
		utils.ResNormalErr(c, 300, "请先登录")
		return false
	}
	return true
}

