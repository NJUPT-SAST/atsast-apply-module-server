package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	
	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/logger"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type UpdateUserSastProfileRequest struct {
	RequesterId *uuid.UUID `json:"requesterId"`
	UserId      *uuid.UUID `json:"userId"`
	*entity.UserSastProfile
}

func UpdateUserSastProfileRequestParser(c *gin.Context) (*UpdateUserSastProfileRequest, error) {
	request := UpdateUserSastProfileRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		return nil, err
	}
	identity := jwt.MustExtractIdentity(c)
	userId := uuid.MustParse(c.Param("userId"))
	request.RequesterId = identity.Uid
	request.UserId = &userId
	logger.LogRequest("UpdateUserSastProfile", request)
	return &request, nil
}

func UpdateUserSastProfileRequestHandler(request *UpdateUserSastProfileRequest) *response.Response {
	isSuperAdmin, err := service.IsSuperAdmin(request.RequesterId)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	if !isSuperAdmin {
		return response.Failed().SetMsg(service.PermissionDeniedErr.Error())
	}

	err = service.UpdateUserSastProfile(request.UserId, request.UserSastProfile)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	return response.Success()
}
