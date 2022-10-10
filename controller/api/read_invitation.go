package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type ReadInvitationRequest struct {
	RequesterId *uuid.UUID `json:"requesterId"`
	Code        *string    `json:"code"`
}

func ReadInvitationRequestParser(c *gin.Context) (*ReadInvitationRequest, error) {
	request := ReadInvitationRequest{}
	identity := jwt.MustExtractIdentity(c)
	code := c.Query("code")
	if code == "" {
		return nil, errors.New("invitation code required")
	}

	request.RequesterId = identity.Uid
	request.Code = &code
	return &request, nil
}

func ReadInvitationRequestHandler(request *ReadInvitationRequest) *response.Response {
	invitation, err := service.ReadInvitation(request.Code)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}

	if invitation.UserId == nil {
		if err = service.UpdateInvitationUserId(request.Code, request.RequesterId); err != nil {
			return response.Failed().SetMsg(err.Error())
		}
		if *invitation.Type == entity.AdminUser {
			if err = service.UpdateUserRole(request.RequesterId, &entity.AdminUser); err != nil {
				return response.Failed().SetMsg(err.Error())
			}
		}
	} else if *invitation.UserId != *request.RequesterId {
		return response.Failed().SetMsg(service.PermissionDeniedErr.Error())
	}

	return response.Success().SetData(invitation)
}
