package controller

import (
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/common/wechat"
	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type loginRequest struct {
	WeChatCode *string `json:"wechatCode" binding:"required"`
}

type loginResponse struct {
	UserId *uuid.UUID `json:"userId" binding:"required"`
	Token  *string    `json:"token" binding:"required"`
}

func login(request *loginRequest) *stdrsp.Result {
	result, err := wechat.Code2Session(request.WeChatCode)
	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallWechatApiErrMsg)
	}
	userId, err := service.ReadUserIdWithCreateIfNotExist(&result.OpenID)
	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
	}

	identity := &jwt.Identity{
		Uid: userId,
	}
	token, err := jwt.NewIdentityJwtString(identity)
	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg("generate token error" + stdrsp.ErrMsgSuffix)
	}

	return stdrsp.Response(stdrsp.Success).Data(loginResponse{
		UserId: userId,
		Token:  token,
	})
}
