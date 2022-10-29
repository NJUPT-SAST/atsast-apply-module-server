package main

import (
	"github.com/gin-gonic/gin"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller"
	"github.com/njupt-sast/atsast-apply-module-server/middleware"
)

func main() {
	r := gin.Default()

	apiRouter := r.Group("api")
	apiRouter.GET("health", controller.CheckHealth)
	apiRouter.GET("config", controller.ReadConfig)
	apiRouter.GET("exam", controller.ReadExamList)
	apiRouter.POST("login", controller.Login)

	invitationRouter := apiRouter.Group("invitation")
	invitationRouter.Use(middleware.BearerTokenAuth[*jwt.Identity](jwt.ParseIdentityJwtString, jwt.InjectIdentity))
	invitationRouter.GET("", controller.ReadInvitation)

	userRouter := apiRouter.Group("user")
	userRouter.Use(middleware.BearerTokenAuth[*jwt.Identity](jwt.ParseIdentityJwtString, jwt.InjectIdentity))
	userRouter.GET("", controller.ReadUser)
	userRouter.GET(":userId/profile", controller.ReadUserProfile)
	userRouter.PUT(":userId/profile", controller.UpdateUserProfile)
	userRouter.GET(":userId/profile/sast", controller.ReadUserSastProfile)
	userRouter.PUT(":userId/profile/sast", controller.UpdateUserSastProfile)
	userRouter.GET(":userId/score", controller.ReadUserScore)
	userRouter.PUT(":userId/score", controller.UpdateUserScore)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
