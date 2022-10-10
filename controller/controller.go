package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/njupt-sast/atsast-apply-module-server/controller/api"
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
)

var (
	CheckHealth       = NewPureController(api.CheckHealth)
	ReadConfig        = NewPureController(api.ReadConfig)
	ReadExamList      = NewPureController(api.ReadExamList)
	Login             = NewController[*api.LoginRequest](api.LoginRequestParser, api.LoginRequestHandler)
	ReadInvitation    = NewController[*api.ReadInvitationRequest](api.ReadInvitationRequestParser, api.ReadInvitationRequestHandler)
	ReadUser          = NewController[*api.ReadUserRequest](api.ReadUserRequestParser, api.ReadUserRequestHandler)
	ReadUserProfile   = NewController[*api.ReadUserProfileRequest](api.ReadUserProfileRequestParser, api.ReadUserProfileRequestHandler)
	UpdateUserProfile = NewController[*api.UpdateUserProfileRequest](api.UpdateUserProfileRequestParser, api.UpdateUserProfileRequestHandler)
	ReadUserScore     = NewController[*api.ReadUserScoreRequest](api.ReadUserScoreRequestParser, api.ReadUserScoreRequestHandler)
	UpdateUserScore   = NewController[*api.UpdateUserScoreRequest](api.UpdateUserScoreRequestParser, api.UpdateUserScoreRequestHandler)
)

type PureRequestHandlerFunc func() *response.Response

func NewPureController(requestHandler PureRequestHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, requestHandler())
	}
}

type RequestParserFunc[T any] func(*gin.Context) (T, error)

type RequestHandlerFunc[T any] func(T) *response.Response

func NewController[T any](
	requestParser RequestParserFunc[T],
	requestHandler RequestHandlerFunc[T],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		request, err := requestParser(c)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, requestHandler(request))
	}
}
