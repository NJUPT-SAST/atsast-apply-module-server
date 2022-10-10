package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId   *uuid.UUID   `json:"userId" bson:"userId,omitempty" binding:"required"`
	WeChatId *string      `json:"weChatId" bson:"weChatId,omitempty"`
	Profile  *UserProfile `json:"profile" bson:"profile,omitempty"`
	ExamMap  *UserExamMap `json:"examMap" bson:"scoreMap,omitempty"`
	Role     *UserRole    `json:"role" bson:"role,omitempty"`
}

type UserProfile struct {
	Name    *string             `json:"name" bson:"name,omitempty" binding:"required"`
	School  *UserProfileSchool  `json:"school" bson:"school,omitempty" binding:"required"`
	Contact *UserProfileContact `json:"contact" bson:"contact,omitempty" binding:"required"`
	Apply   *UserProfileApply   `json:"apply" bson:"apply,omitempty" binding:"required"`
}

type UserProfileSchool struct {
	StudentId *string `json:"studentId" bson:"studentId,omitempty" binding:"required"`
	College   *string `json:"college" bson:"college,omitempty" binding:"required"`
	Major     *string `json:"major" bson:"major,omitempty" binding:"required"`
}

type UserProfileContact struct {
	Phone *string `json:"phone" bson:"phone,omitempty" binding:"required"`
	QQ    *string `json:"qq" bson:"qq,omitempty" binding:"required"`
}

type UserProfileApply struct {
	Choice1 *string `json:"choice1" bson:"choice1,omitempty" binding:"required"`
	Choice2 *string `json:"choice2" bson:"choice2,omitempty" binding:"required"`
}

type UserExamMap map[string]UserScoreMap

type UserScoreMap map[string]UserScore

type UserScore struct {
	Score     *int       `json:"score" bson:"score,omitempty" binding:"required"`
	JudgerId  *uuid.UUID `json:"judger" bson:"judger,omitempty" binding:"required"`
	JudgeTime *time.Time `json:"judgeTime" bson:"judgeTime,omitempty" binding:"required"`
}

type UserRole string

func (userRole *UserRole) IsAdmin() bool {
	return userRole != nil && (*userRole == AdminUser || *userRole == SuperAdminUser)
}

var (
	CommonUser     UserRole = "user"
	AdminUser      UserRole = "admin"
	SuperAdminUser UserRole = "super admin"
)
