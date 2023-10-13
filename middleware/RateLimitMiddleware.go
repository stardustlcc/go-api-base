package middleware

import (
	"go-api-base/pkg/app"
	"go-api-base/pkg/errcode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(fillInterval time.Duration, capacity int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			appRes := app.NewApp(ctx)
			appRes.ErrorJson(http.StatusOK, errcode.RATE_LIMIT)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
