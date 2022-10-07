package dao

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/njupt-sast/atsast-apply-module-server/entity"
)

func ReadInvitation(code *string) (*entity.Invitation, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	var invitation entity.Invitation
	err := InvitationColl.FindOne(ctx, bson.D{{Key: "code", Value: code}}).Decode(&invitation)
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func UpdateInvitationUserId(code *string, userId *uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	result := InvitationColl.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "code", Value: code}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "userId", Value: userId},
		}}},
	)
	return result.Err()
}
