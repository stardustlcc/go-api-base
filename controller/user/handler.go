package user

import (
	"go-api-base/initialize/mysql"
	"go-api-base/initialize/redis"
	"go-api-base/services/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	CreateUser(c *gin.Context) //创建用户
	UserInfo(c *gin.Context)   //查询用户信息
}

type handler struct {
	logger      *zap.Logger
	db          mysql.Repo
	cache       redis.Repo
	userService user.Service
}

// 通过New创建 handler 结构体
func NewController(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		db:          db,
		cache:       cache,
		userService: user.NewService(db, cache),
	}
}
