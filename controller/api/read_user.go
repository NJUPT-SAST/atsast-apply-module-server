package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type ReadUserRequest struct {
	RequesterId *uuid.UUID `json:"requesterId"`
	StudentId   *string    `json:"studentId"`
}

func ReadUserRequestParser(c *gin.Context) (*ReadUserRequest, error) {
	request := ReadUserRequest{}

	identity := jwt.MustExtractIdentity(c)
	studentId := c.Query("studentId")

	request.RequesterId = identity.Uid
	request.StudentId = &studentId
	return &request, nil
}

type ReadUserResponse struct {
	UserList []*uuid.UUID `json:"userList" binding:"required"`
}

func ReadUserRequestHandler(request *ReadUserRequest) *response.Response {
	isAdmin, err := service.IsAdmin(request.RequesterId)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	if !isAdmin {
		return response.Failed().SetMsg(service.PermissionDeniedErr.Error())
	}

	userList, err := service.ReadUserBySpecifyProfileField("school.studentId", *request.StudentId)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	userIdList := make([]*uuid.UUID, 0)
	for _, user := range userList {
		userIdList = append(userIdList, user.UserId)
	}
	return response.Success().SetData(ReadUserResponse{
		UserList: userIdList,
	})
}
