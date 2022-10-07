package service

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/njupt-sast/atsast-apply-module-server/dao"
	"github.com/njupt-sast/atsast-apply-module-server/entity"
)

func ReadInvitation(code *string) (*entity.Invitation, error) {
	invitation, err := dao.ReadInvitation(code)
	if err == mongo.ErrNoDocuments {
		return nil, DocumentNotFoundError
	}

	if err != nil {
		return nil, err
	}

	return invitation, nil
}

func UpdateInvitationUserId(code *string, userId *uuid.UUID) error {
	return dao.UpdateInvitationUserId(code, userId)
}
