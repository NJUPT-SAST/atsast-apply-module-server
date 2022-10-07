package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileSchool struct {
	StudentId *string `json:"studentId" bson:"studentId" binding:"required"`
	College   *string `json:"college" bson:"college" binding:"required"`
	Major     *string `json:"major" bson:"major" binding:"required"`
}

type UserProfileContact struct {
	Phone *string `json:"phone" bson:"phone" binding:"required"`
	QQ    *string `json:"qq" bson:"qq" binding:"required"`
}

type UserProfileApply struct {
	Choice1 *string `json:"choice1" bson:"choice1" binding:"required"`
	Choice2 *string `json:"choice2" bson:"choice2" binding:"required"`
}

type UserProfile struct {
	Name    *string             `json:"name" bson:"name" binding:"required"`
	School  *UserProfileSchool  `json:"school" bson:"school" binding:"required"`
	Contact *UserProfileContact `json:"contact" bson:"contact" binding:"required"`
	Apply   *UserProfileApply   `json:"apply" bson:"apply" binding:"required"`
}

type UserScore struct {
	Score     *int       `json:"score" bson:"score" binding:"required"`
	JudgerId  *uuid.UUID `json:"judger" bson:"judger" binding:"required"`
	JudgeTime *time.Time `json:"judgeTime" bson:"judgeTime" binding:"required"`
}

type UserScoreMap map[string]UserScore

type UserExamMap map[string]UserScoreMap

type UserRole int

var (
	GeneralUserRole UserRole = 1
	AdminUserRole   UserRole = 2
)

type User struct {
	UserId   *uuid.UUID   `json:"userId" bson:"userId" binding:"required"`
	WeChatId *string      `json:"weChatId" bson:"weChatId" binding:"required"`
	Profile  *UserProfile `json:"profile" bson:"profile" binding:"required"`
	ExamMap  *UserExamMap `json:"examMap" bson:"scoreMap" binding:"required"`
	Role     *UserRole    `json:"role" bson:"role" binding:"required"`
}
