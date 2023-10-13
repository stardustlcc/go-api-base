package main

import (
	"context"
	"go-api-base/initialize"
	"go-api-base/pkg/global"
	"go-api-base/pkg/logger"
	"go-api-base/pkg/shutdown"
	"go-api-base/pkg/util/timeutil"
	"go-api-base/router"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func init() {
	initialize.NewConfig()
}

func main() {

	accessLogger, err := logger.NewLogger(
		logger.WithLogPath(global.Conf.Logs.Path),
		logger.WithTimeLayout(timeutil.CSTLayout),
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
	}()

	s, err := router.NewRouter(accessLogger)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:           ":" + global.Conf.Server.HttpPort,
		Handler:        s.Router,
		ReadTimeout:    global.Conf.Server.ReadTimeOut * time.Second,
		WriteTimeout:   global.Conf.Server.WriteTimeOut * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//开启goroutine启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	//优雅关机
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},
		func() {
			if s.Db != nil {
				if err := s.Db.DBClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}
			}
		},
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},
	)
}
