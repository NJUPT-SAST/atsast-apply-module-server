package controller

import (
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
	"github.com/njupt-sast/atsast-apply-module-server/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type readUserProfileResponse struct {
	UserId  *uuid.UUID          `json:"userId"`
	Profile *entity.UserProfile `json:"profile"`
}

func readUserProfile(identity *jwt.Identity, requestUserId *uuid.UUID) *stdrsp.Result {
	if *identity.Uid != *requestUserId {
		user, err := service.ReadUser(identity.Uid)
		if err == service.DocumentNotFoundError {
			return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
		}

		if err != nil {
			return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
		}

		if !service.IsAdmin(user.Role) {
			return stdrsp.Response(stdrsp.BadRequest).Msg(stdrsp.PermissionDeniedErrMsg)
		}
	}

	user, err := service.ReadUser(requestUserId)
	if err == service.DocumentNotFoundError {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
	}

	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
	}

	return stdrsp.Response(stdrsp.Success).Data(readUserProfileResponse{
		UserId:  requestUserId,
		Profile: user.Profile,
	})
}

type updateUserProfileRequest struct {
	Name   *string `json:"name" binding:"required"`
	School *struct {
		StudentId *string `json:"studentId" binding:"required"`
		College   *string `json:"college" binding:"required"`
		Major     *string `json:"major" binding:"required"`
	} `json:"school" binding:"required"`
	Contact *struct {
		Phone *string `json:"phone" binding:"required"`
		QQ    *string `json:"qq" binding:"required"`
	} `json:"contact" binding:"required"`
	Apply *struct {
		Choice1 *string `json:"choice1" binding:"required"`
		Choice2 *string `json:"choice2"`
	} `json:"apply" binding:"required"`
}

func (request *updateUserProfileRequest) UserProfile() *entity.UserProfile {
	return &entity.UserProfile{
		Name: request.Name,
		School: &entity.UserProfileSchool{
			StudentId: request.School.StudentId,
			College:   request.School.College,
			Major:     request.School.Major,
		},
		Contact: &entity.UserProfileContact{
			Phone: request.Contact.Phone,
			QQ:    request.Contact.QQ,
		},
		Apply: &entity.UserProfileApply{
			Choice1: request.Apply.Choice1,
			Choice2: request.Apply.Choice2,
		},
	}
}

func updateUserProfile(identity *jwt.Identity, requestUserId *uuid.UUID, request *updateUserProfileRequest) *stdrsp.Result {
	if *identity.Uid != *requestUserId {
		return stdrsp.Response(stdrsp.BadRequest).Msg(stdrsp.PermissionDeniedErrMsg)
	}

	if err := service.UpdateUserProfile(requestUserId, request.UserProfile()); err == service.DocumentNotFoundError {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
	} else if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
	}

	return stdrsp.Response(stdrsp.Success)
}
