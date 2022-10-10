package service

import (
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/model/dao"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
)

func IsAdmin(userId *uuid.UUID) (bool, error) {
	user, err := dao.ReadUser(userId)
	if err == dao.NoDocumentsErr {
		return false, NotFoundErr
	}
	if err != nil {
		return false, CallDatabaseErr
	}
	return user.Role.IsAdmin(), nil
}

func ReadUserWithCreateIfNotExist(weChatId *string) (*entity.User, error) {
	return dao.ReadUserWithCreateIfNotExist(weChatId)
}

func ReadUser(userId *uuid.UUID) (*entity.User, error) {
	user, err := dao.ReadUser(userId)
	if err == dao.NoDocumentsErr {
		return nil, NotFoundErr
	}
	if err != nil {
		return nil, CallDatabaseErr
	}
	return user, nil
}

func ReadUserBySpecifyProfileField(fieldName string, fieldValue interface{}) ([]entity.User, error) {
	userList, err := dao.ReadUserListBySpecifyField("profile."+fieldName, fieldValue)
	if err == dao.NoDocumentsErr {
		return nil, NotFoundErr
	}
	if err != nil {
		return nil, CallDatabaseErr
	}
	return userList, nil
}

func UpdateUserRole(userId *uuid.UUID, userRole *entity.UserRole) error {
	err := dao.UpdateUserRole(userId, userRole)
	if err == dao.NoDocumentsErr {
		return NotFoundErr
	}
	if err != nil {
		return CallDatabaseErr
	}
	return nil

}

func UpdateUserProfile(userId *uuid.UUID, userProfile *entity.UserProfile) error {
	err := dao.UpdateUserProfile(userId, userProfile)
	if err == dao.NoDocumentsErr {
		return NotFoundErr
	}
	if err != nil {
		return CallDatabaseErr
	}
	return nil
}

func UpdateUserScore(userId *uuid.UUID, examId *string, userScoreMap *entity.UserScoreMap) error {
	err := dao.UpdateUserScore(userId, examId, userScoreMap)
	if err == dao.NoDocumentsErr {
		return NotFoundErr
	}
	if err != nil {
		return CallDatabaseErr
	}
	return nil
}
