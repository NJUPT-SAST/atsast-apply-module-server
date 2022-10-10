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

type UpdateUserProfileRequest struct {
	RequesterId *uuid.UUID `json:"requesterId"`
	UserId      *uuid.UUID `json:"userId"`
	*entity.UserProfile
}

func UpdateUserProfileRequestParser(c *gin.Context) (*UpdateUserProfileRequest, error) {
	request := UpdateUserProfileRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		return nil, err
	}
	identity := jwt.MustExtractIdentity(c)
	userId := uuid.MustParse(c.Param("userId"))
	request.RequesterId = identity.Uid
	request.UserId = &userId
	logger.LogRequest("UpdateUserProfile", request)
	return &request, nil
}

func UpdateUserProfileRequestHandler(request *UpdateUserProfileRequest) *response.Response {
	if *request.RequesterId != *request.UserId {
		isAdmin, err := service.IsAdmin(request.RequesterId)
		if err != nil {
			return response.Failed().SetMsg(err.Error())
		}
		if !isAdmin {
			return response.Failed().SetMsg(service.PermissionDeniedErr.Error())
		}
	}

	err := service.UpdateUserProfile(request.UserId, request.UserProfile)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	return response.Success()
}
