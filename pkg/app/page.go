package app

import (
	"go-api-base/pkg/global"
	"go-api-base/pkg/util/stringutil"

	"github.com/gin-gonic/gin"
)

func GetPageNum(ctx *gin.Context) int {
	pageNum := stringutil.StrTo(ctx.Query("pageNum")).MustInt()
	if pageNum <= 0 {
		return 1
	}
	return pageNum
}

func GetPageLimit(ctx *gin.Context) int {
	pageLimit := stringutil.StrTo(ctx.Query("pageLimit")).MustInt()
	if pageLimit <= 0 {
		return global.Conf.App.DefaultPageSize
	}
	if pageLimit > global.Conf.App.MaxPageSize {
		return global.Conf.App.MaxPageSize
	}
	return pageLimit
}

func GetPageOffset(pageNum, pageLimit int) int {
	result := 0
	if pageNum > 0 {
		result = (pageNum - 1) * pageLimit
	}
	return result
}
