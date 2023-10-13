package user

import (
	"go-api-base/pkg/app"

	"github.com/gin-gonic/gin"
)

type userInfoRequest struct {
	UserId int16 `form:"username"`
}

type userResponse struct {
	UserName string `json:"username"`
	UserId   int16  `josn:"id"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

func (h *handler) UserInfo(ctx *gin.Context) {

	_ = new(userInfoRequest)
	_ = new(userResponse)
	username, id, age, address := h.userService.UserInfo(ctx)

	app := app.NewApp(ctx)
	app.SuccessJson(userResponse{
		UserName: username,
		UserId:   id,
		Age:      age,
		Address:  address,
	})

}
