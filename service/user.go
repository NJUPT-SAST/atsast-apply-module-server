package service

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/njupt-sast/atsast-apply-module-server/dao"
	"github.com/njupt-sast/atsast-apply-module-server/entity"
)

func ReadUserIdWithCreateIfNotExist(weChatId *string) (*uuid.UUID, error) {
	user, err := dao.FindUserWithCreateIfNotExist(weChatId)
	if err != nil {
		return nil, err
	}

	return user.UserId, err
}

func IsAdmin(userRole *entity.UserRole) bool {
	return userRole != nil && *userRole == entity.AdminUserRole
}

func ReadUser(userId *uuid.UUID) (*entity.User, error) {
	user, err := dao.ReadUser(userId)
	if err == mongo.ErrNoDocuments {
		return nil, DocumentNotFoundError
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserRole(userId *uuid.UUID, userRole *entity.UserRole) error {
	err := dao.UpdateUserRole(userId, userRole)
	if err == mongo.ErrNoDocuments {
		return DocumentNotFoundError
	}

	return err
}

func UpdateUserProfile(userId *uuid.UUID, userProfile *entity.UserProfile) error {
	err := dao.UpdateUserProfile(userId, userProfile)
	if err == mongo.ErrNoDocuments {
		return DocumentNotFoundError
	}

	return err
}

func UpdateUserScore(userId *uuid.UUID, examId *string, userScoreMap *entity.UserScoreMap) error {
	err := dao.UpdateUserScore(userId, examId, userScoreMap)
	if err == mongo.ErrNoDocuments {
		return DocumentNotFoundError
	}

	return err
}
