package api

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/logger"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type UpdateUserScoreRequestOrigin struct {
	ExamId    *string `json:"examId" binding:"required"`
	ScoreList []struct {
		ProblemId *string `json:"problemId" binding:"required"`
		Score     *int    `json:"score" binding:"required"`
	} `json:"scoreList" binding:"required"`
}

func (request *UpdateUserScoreRequestOrigin) UserScoreMap(judgerId *uuid.UUID) (*entity.UserScoreMap, error) {
	userScoreMap := entity.UserScoreMap{}
	currentTime := time.Now()
	if strings.Contains(*request.ExamId, ".") {
		return nil, errors.New("`.` in problem id is not allowed")
	}
	for _, scores := range request.ScoreList {
		if strings.Contains(*scores.ProblemId, ".") {
			return nil, errors.New("`.` in problem id is not allowed")
		}
		userScoreMap[*scores.ProblemId] = entity.UserScore{
			Score:     scores.Score,
			JudgerId:  judgerId,
			JudgeTime: &currentTime,
		}
	}
	return &userScoreMap, nil
}

type UpdateUserScoreRequest struct {
	RequesterId  *uuid.UUID           `json:"requesterId"`
	UserId       *uuid.UUID           `json:"userId"`
	ExamId       *string              `json:"examId"`
	UserScoreMap *entity.UserScoreMap `json:"userScoreMap"`
}

func UpdateUserScoreRequestParser(c *gin.Context) (*UpdateUserScoreRequest, error) {
	requestOrigin := UpdateUserScoreRequestOrigin{}
	if err := c.BindJSON(&requestOrigin); err != nil {
		return nil, err
	}
	request := UpdateUserScoreRequest{}
	identity := jwt.MustExtractIdentity(c)
	userScoreMap, err := requestOrigin.UserScoreMap(identity.Uid)
	if err != nil {
		return nil, err
	}
	request.UserScoreMap = userScoreMap
	userId := uuid.MustParse(c.Param("userId"))
	request.RequesterId = identity.Uid
	request.UserId = &userId
	request.ExamId = requestOrigin.ExamId
	logger.LogRequest("UpdateUserScore", request)
	return &request, nil
}

func UpdateUserScoreRequestHandler(request *UpdateUserScoreRequest) *response.Response {
	isAdmin, err := service.IsAdmin(request.RequesterId)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	if !isAdmin {
		return response.Failed().SetMsg(service.PermissionDeniedErr.Error())
	}

	err = service.UpdateUserScore(request.UserId, request.ExamId, request.UserScoreMap)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}

	return response.Success()
}
