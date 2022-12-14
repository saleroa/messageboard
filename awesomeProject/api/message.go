// api层除了router.go以外的包都是不同的功能块
//拿serve层的函数

package api

//message的内容用myslq存储。
//现在是识别每个用户的username
import (
	"awesomeProject/middleware"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/utils"

	"github.com/gin-gonic/gin"
)

//拿到该用户的username

func GetMessage(c *gin.Context) {
	if middleware.Auth(c) {
		service.GainMessage(User) //username
		utils.RespOk(c)
	}
	//utils.ResNormalErr(c,500,"请先登录")
}

func SendMessage(c *gin.Context) {
	message := c.PostForm("message")

	if middleware.Auth(c) {
		service.AddMessage(model.Message{
			Username: User,
			Message:  message, /////////////////////////////////////////
		})
		utils.RespOk(c)
	}
	//utils.ResNormalErr(c,500,"请先登录")

}

func PutMessage(c *gin.Context) {
	message := c.PostForm("message")

	if middleware.Auth(c) {
		service.FixMessage(User, message) //username ,
		utils.RespOk(c)
	}
	//utils.ResNormalErr(c,500,"请先登录")
}

func DelMessage(c *gin.Context) {

	if middleware.Auth(c) {
		service.SubMessage(User) //username
		utils.RespOk(c)
	}
	//utils.ResNormalErr(c,500,"请先登录")

}
