package api

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

type ReadUserScoreRequest struct {
	RequesterId *uuid.UUID `json:"requesterId"`
	UserId      *uuid.UUID `json:"userId"`
	ExamId      *string    `json:"examId"`
}

func ReadUserScoreRequestParser(c *gin.Context) (*ReadUserScoreRequest, error) {
	identity := jwt.MustExtractIdentity(c)
	examId := c.Query("examId")
	if examId == "" {
		return nil, errors.New("exam id required")
	}
	userId := uuid.MustParse(c.Param("userId"))

	request := ReadUserScoreRequest{}
	request.RequesterId = identity.Uid
	request.UserId = &userId
	request.ExamId = &examId
	return &request, nil
}

func ReadUserScoreRequestHandler(request *ReadUserScoreRequest) *response.Response {
	if *request.RequesterId != *request.UserId {
		isAdmin, err := service.IsAdmin(request.RequesterId)
		if err != nil {
			return response.Failed().SetMsg(err.Error())
		}
		if !isAdmin {
			return response.Failed().SetMsg(service.PermissionDeniedErr.Error())
		}
	}

	user, err := service.ReadUser(request.UserId)
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	return response.Success().SetData(newReadUserScoreResponse(request.RequesterId, request.ExamId, user.ExamMap))
}

type userScoreResponse struct {
	ProblemId string     `json:"problemId" bson:"problemId,omitempty" binding:"required"`
	Score     *int       `json:"score" bson:"score,omitempty" binding:"required"`
	JudgerId  *uuid.UUID `json:"judger" bson:"judger,omitempty" binding:"required"`
	JudgeTime *time.Time `json:"judgeTime" bson:"judgeTime,omitempty" binding:"required"`
}

type ReadUserScoreResponse struct {
	UserId    *uuid.UUID          `json:"userId" binding:"required"`
	ExamId    *string             `json:"examId" binding:"required"`
	ScoreList []userScoreResponse `json:"scoreList" binding:"required"`
}

func newReadUserScoreResponse(userId *uuid.UUID, examId *string, userExamMap *entity.UserExamMap) *ReadUserScoreResponse {
	scoreList := make([]userScoreResponse, 0)

	if userExamMap != nil {
		userScoreMap := (*userExamMap)[*examId]
		for problemId, userScore := range userScoreMap {
			scoreList = append(scoreList, userScoreResponse{
				ProblemId: problemId,
				Score:     userScore.Score,
				JudgerId:  userScore.JudgerId,
				JudgeTime: userScore.JudgeTime,
			})
		}
	}
	return &ReadUserScoreResponse{
		UserId:    userId,
		ExamId:    examId,
		ScoreList: scoreList,
	}
}
