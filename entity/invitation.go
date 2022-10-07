package entity

import "github.com/google/uuid"

type Invitation struct {
	Code    *string     `json:"code" bson:"code,omitempty" binding:"required"`
	Type    *string     `json:"type" bson:"type,omitempty" binding:"required"`
	Profile *PreProfile `json:"profile" bson:"profile,omitempty" binding:"required"`
	UserId  *uuid.UUID  `json:"userId" bson:"userId,omitempty" binding:"required"`
}

type PreProfile struct {
	Name      *string `json:"name" bson:"name,omitempty" binding:"required"`
	StudentId *string `json:"studentId" bson:"studentId,omitempty" binding:"required"`
	College   *string `json:"college" bson:"college,omitempty" binding:"required"`
	Major     *string `json:"major" bson:"major,omitempty" binding:"required"`
	Choice1   *string `json:"choice1" bson:"choice1,omitempty" binding:"required"`
	Choice2   *string `json:"choice2" bson:"choice2,omitempty" binding:"required"`
	Phone     *string `json:"phone" bson:"phone,omitempty" binding:"required"`
	QQ        *string `json:"qq" bson:"qq,omitempty" binding:"required"`
}
