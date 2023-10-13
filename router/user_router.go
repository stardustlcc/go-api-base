package router

import (
	"go-api-base/controller/user"
)

func setUserRouter(r *resource) {

	user := user.NewController(r.logger, r.db, r.cache)

	router := r.router.Group("/user")
	{
		router.GET("/userInfo", user.UserInfo)
	}
}
