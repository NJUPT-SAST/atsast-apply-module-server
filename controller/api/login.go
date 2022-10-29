package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/common/wechat"
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type LoginRequest struct {
	WeChatCode *string `json:"wechatCode" binding:"required"`
}

type LoginResponse struct {
	UserId *uuid.UUID       `json:"userId" binding:"required"`
	Token  *string          `json:"token" binding:"required"`
	Role   *entity.UserRole `json:"role" binding:"required"`
}

func LoginRequestParser(c *gin.Context) (*LoginRequest, error) {
	request := LoginRequest{}
	if err := c.BindJSON(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func LoginRequestHandler(request *LoginRequest) *response.Response {
	result, err := wechat.Code2Session(request.WeChatCode)
	if err != nil {
		return response.Failed()
	}
	user, err := service.ReadUserWithCreateIfNotExist(&result.OpenID)
	if err != nil {
		return response.Failed()
	}

	token, err := jwt.NewIdentityString(&jwt.Identity{
		Uid: user.UserId,
	})
	if err != nil {
		return response.Failed()
	}

	return response.Success().SetData(LoginResponse{
		UserId: user.UserId,
		Role:   user.Role,
		Token:  token,
	})
}
