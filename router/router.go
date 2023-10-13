package router

import (
	"go-api-base/initialize/mysql"
	"go-api-base/initialize/redis"
	"go-api-base/middleware"
	"go-api-base/pkg/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Db     mysql.Repo
	Cache  redis.Repo
	Router *gin.Engine
}

type resource struct {
	logger *zap.Logger
	db     mysql.Repo
	cache  redis.Repo
	router *gin.Engine
}

func NewRouter(logger *zap.Logger) (*Server, error) {

	r := new(resource)
	r.logger = logger

	db, err := mysql.NewMysql()
	if err != nil {
		logger.Fatal("new db is error", zap.Error(err))
	}
	r.db = db

	rdb, err := redis.NewRedis()
	if err != nil {
		logger.Fatal("new redis is error", zap.Error(err))
	}
	r.cache = rdb

	//设置模式
	gin.SetMode(global.Conf.Server.RunModel)

	//创建不带中间件的路由
	r.router = gin.New()

	//设置中间件
	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())
	r.router.Use(middleware.CORSMiddleware())

	//注册路由
	setUserRouter(r)

	s := new(Server)
	s.Db = r.db
	s.Cache = r.cache
	s.Router = r.router
	return s, nil
}
