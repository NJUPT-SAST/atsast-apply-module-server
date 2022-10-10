package service

import (
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/model/dao"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
)

func ReadInvitation(code *string) (*entity.Invitation, error) {
	invitation, err := dao.ReadInvitation(code)
	if err == dao.NoDocumentsErr {
		return nil, NotFoundErr
	}
	if err != nil {
		return nil, CallDatabaseErr
	}
	return invitation, nil
}

func UpdateInvitationUserId(code *string, userId *uuid.UUID) error {
	err := dao.UpdateInvitationUserId(code, userId)
	if err == dao.NoDocumentsErr {
		return NotFoundErr
	}
	if err != nil {
		return CallDatabaseErr
	}
	return nil
}
