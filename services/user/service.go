package user

import (
	"go-api-base/initialize/mysql"
	"go-api-base/initialize/redis"

	"github.com/gin-gonic/gin"
)

var _ Service = (*service)(nil)

type Service interface {
	UserInfo(ctx *gin.Context) (usernam string, id int16, age int, address string)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func NewService(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}
