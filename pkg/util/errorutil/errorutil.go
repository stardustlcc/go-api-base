package errorutil

import (
	"fmt"
	"go-api-base/pkg/app"
	"go-api-base/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 捕捉错误
func RecoverErr(ctx *gin.Context, logger *zap.Logger) {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("%+v", err)
		logger.Error("err msg: " + msg)
		response := app.NewApp(ctx)
		response.ErrorJson(http.StatusOK, errcode.SYSTEM_ERROR)
	}
}

// 检测错误
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
