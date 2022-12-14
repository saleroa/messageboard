//cmd包仅仅是用来装main文件的

package main

import (
	"awesomeProject/api"
	"awesomeProject/dao"
)

func main() {
	dao.Initdb()
	api.Initrouter()
}
