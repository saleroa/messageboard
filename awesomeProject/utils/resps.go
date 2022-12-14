package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建一个回复的模板
type respTemplate struct {
	Stasus int    `json:"status"`
	Info   string `json:"info"`
}

var OK = respTemplate{
	Stasus: 200,
	Info:   "success",
}

var ParamError = respTemplate{ //参数错误
	Stasus: 300,
	Info:   "param error",
}

var internalError = respTemplate{ //服务器内部错误
	Stasus: 500,
	Info:   "internal error",
}

func RespOk(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
}

func RespParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

func ResInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, internalError)
}

func ResNormalErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
