package main

import (
	"github.com/gin-gonic/gin"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller"
	"github.com/njupt-sast/atsast-apply-module-server/middleware"
)

func main() {
	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	apiRouter := r.Group("api")

	apiRouter.GET("health", controller.Health)
	apiRouter.POST("login", controller.Login)

	userRouter := apiRouter.Group("user")
	userRouter.Use(middleware.BearerAuth(jwt.ParseIdentityJwtString, jwt.InjectIdentity))
	userRouter.GET(":userId/profile", controller.ReadUserProfile)
	userRouter.PUT(":userId/profile", controller.UpdateUserProfile)
	userRouter.GET(":userId/score", controller.ReadUserScore)
	userRouter.PUT(":userId/score", controller.UpdateUserScore)

	examRouter := apiRouter.Group("exam")
	examRouter.GET("", controller.ReadExamList)

	configRouter := apiRouter.Group("config")
	configRouter.GET("", controller.ReadConfig)

	invitationRouter := apiRouter.Group("invitation")
	invitationRouter.Use(middleware.BearerAuth(jwt.ParseIdentityJwtString, jwt.InjectIdentity))
	invitationRouter.GET("", controller.ReadInvitation)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
