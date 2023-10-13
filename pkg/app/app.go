package app

import (
	"go-api-base/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ App = (*app)(nil)

type App interface {
	SuccessJson(data interface{})
	ErrorJson(httpCode, errCode int)
}

type app struct {
	ctx  *gin.Context
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewApp(ctx *gin.Context) App {
	return &app{
		ctx: ctx,
	}
}

func (a *app) SuccessJson(data interface{}) {
	a.ctx.JSON(http.StatusOK, app{
		Code: errcode.SUCCESS,
		Msg:  errcode.GetMsg(errcode.SUCCESS),
		Data: data,
	})
}

func (a *app) ErrorJson(httpCode, errCode int) {
	errMsg := errcode.GetMsg(errCode)
	a.ctx.JSON(httpCode, app{
		Code: errCode,
		Msg:  errMsg,
		Data: gin.H{},
	})
}
