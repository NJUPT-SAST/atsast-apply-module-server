package controller

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
	"github.com/njupt-sast/atsast-apply-module-server/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type userScoreResponse struct {
	ProblemId string     `json:"problemId" binding:"required"`
	Score     *int       `json:"score" binding:"required"`
	JudgerId  *uuid.UUID `json:"judger" binding:"required"`
	JudgeTime *time.Time `json:"judgeTime" binding:"required"`
}

type ReadUserScoreResponse struct {
	UserId *uuid.UUID          `json:"userId" binding:"required"`
	Score  []userScoreResponse `json:"score" binding:"required"`
}

func newReadUserScoreResponse(userId *uuid.UUID, userScoreMap *entity.UserScoreMap) *ReadUserScoreResponse {
	response := ReadUserScoreResponse{
		UserId: userId,
	}
	if userScoreMap != nil {
		for problemId, userScore := range *userScoreMap {
			response.Score = append(response.Score, userScoreResponse{
				ProblemId: problemId,
				Score:     userScore.Score,
				JudgerId:  userScore.JudgerId,
				JudgeTime: userScore.JudgeTime,
			})
		}
	}
	return &response
}

func readUserScore(identity *jwt.Identity, examId *string, requestUserId *uuid.UUID) *stdrsp.Result {
	if *identity.Uid != *requestUserId {
		if user, err := service.ReadUser(identity.Uid); err == service.DocumentNotFoundError {
			return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
		} else if err != nil {
			return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
		} else if !service.IsAdmin(user.Role) {
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

	examMap := (*user.ExamMap)[*examId]

	return stdrsp.Response(stdrsp.Success).Data(newReadUserScoreResponse(requestUserId, &examMap))
}

type updateUserScoreRequest struct {
	ExamId    *string `json:"examId" binding:"required"`
	ScoreList []struct {
		ProblemId *string `json:"problemId" binding:"required"`
		Score     *int    `json:"score" binding:"required"`
	} `json:"scoreList" binding:"required"`
}

func (request *updateUserScoreRequest) UserScoreMap(judgerId *uuid.UUID) (*entity.UserScoreMap, error) {
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

func updateUserScore(identity *jwt.Identity, requestUserId *uuid.UUID, request *updateUserScoreRequest) *stdrsp.Result {
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

	userScoreMap, err := request.UserScoreMap(identity.Uid)
	if err != nil {
		return stdrsp.Response(stdrsp.BadRequest).Msg(stdrsp.BadRequestErrMsg)
	}

	err = service.UpdateUserScore(requestUserId, request.ExamId, userScoreMap)
	if err == service.DocumentNotFoundError {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
	}

	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
	}

	return stdrsp.Response(stdrsp.Success)
}
