package user

import "github.com/gin-gonic/gin"

func (s *service) UserInfo(ctx *gin.Context) (usernam string, id int16, age int, address string) {
	usernam = "cici"
	id = 1001
	age = 18
	address = "上海松江区洞泾镇"
	return
}
