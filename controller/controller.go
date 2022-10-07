package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
	"github.com/njupt-sast/atsast-apply-module-server/logger"
)

func responseBindRequestError(c *gin.Context) {
	stdrsp.ResponseResult(c, stdrsp.Response(stdrsp.BadRequest).Msg(stdrsp.BadRequestErrMsg))
}

func Health(c *gin.Context) {
	stdrsp.ResponseResult(c, stdrsp.Response(stdrsp.Success))
}

func ReadConfig(c *gin.Context) {
	stdrsp.ResponseResult(c, readConfig())
}

func Login(c *gin.Context) {
	request := &loginRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		responseBindRequestError(c)
		return
	}

	stdrsp.ResponseResult(c, login(request))
}

func ReadInvitation(c *gin.Context) {
	identity := jwt.ExtractIdentity(c)
	code := c.Query("code")

	stdrsp.ResponseResult(c, readInvitation(identity, &code))
}

func ReadUserProfile(c *gin.Context) {
	identity := jwt.ExtractIdentity(c)
	requestUserId := uuid.MustParse(c.Param("userId"))

	stdrsp.ResponseResult(c, readUserProfile(identity, &requestUserId))
}

func UpdateUserProfile(c *gin.Context) {
	identity := jwt.ExtractIdentity(c)
	requestUserId := uuid.MustParse(c.Param("userId"))

	request := &updateUserProfileRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		responseBindRequestError(c)
		return
	}

	logger.LogRequest("UpdateUserProfile", identity.Uid, request)
	stdrsp.ResponseResult(c, updateUserProfile(identity, &requestUserId, request))
}

func ReadExamList(c *gin.Context) {
	stdrsp.ResponseResult(c, readExamList())
}

func ReadUserScore(c *gin.Context) {
	identity := jwt.ExtractIdentity(c)
	requestUserId := uuid.MustParse(c.Param("userId"))
	examId := c.Query("examId")

	stdrsp.ResponseResult(c, readUserScore(identity, &examId, &requestUserId))
}

func UpdateUserScore(c *gin.Context) {
	identity := jwt.ExtractIdentity(c)
	requestUserId := uuid.MustParse(c.Param("userId"))

	request := &updateUserScoreRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		responseBindRequestError(c)
		return
	}

	logger.LogRequest("UpdateUserScore", identity.Uid, request)
	stdrsp.ResponseResult(c, updateUserScore(identity, &requestUserId, request))
}
