package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
)

func ReadUserWithCreateIfNotExist(weChatId *string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	newUserId := uuid.New()
	newUser := entity.User{
		UserId:   &newUserId,
		WeChatId: weChatId,
		Role:     &entity.CommonUser,
	}

	oldUser := entity.User{}

	err := UserColl.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "weChatId", Value: weChatId}},
		bson.D{{Key: "$setOnInsert", Value: newUser}},
		options.FindOneAndUpdate().SetUpsert(true),
	).Decode(&oldUser)
	if err == mongo.ErrNoDocuments {
		return &newUser, nil
	}
	if err != nil {
		return nil, err
	}
	return &oldUser, nil
}

func ReadUserListBySpecifyField(fieldName string, fieldValue interface{}) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()
	cur, err := UserColl.Find(
		ctx,
		bson.D{{Key: fieldName, Value: fieldValue}},
	)
	if err != nil {
		return nil, err
	}

	userList := make([]entity.User, 0)
	err = cur.All(ctx, &userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func ReadUser(userId *uuid.UUID) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	var user entity.User
	err := UserColl.FindOne(
		ctx,
		bson.D{{Key: "userId", Value: userId}},
	).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserRole(userId *uuid.UUID, userRole *entity.UserRole) error {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	result := UserColl.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "userId", Value: userId}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "role", Value: userRole},
		}}},
	)
	if result.Err() == mongo.ErrNoDocuments {
		return NoDocumentsErr
	}
	return result.Err()
}

func UpdateUserProfile(userId *uuid.UUID, userProfile *entity.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	result := UserColl.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "userId", Value: userId}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "profile", Value: userProfile},
		}}},
	)
	if result.Err() == mongo.ErrNoDocuments {
		return NoDocumentsErr
	}
	return result.Err()
}

func UpdateUserSastProfile(userId *uuid.UUID, userSastProfile *entity.UserSastProfile) error {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	result := UserColl.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "userId", Value: userId}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "sastProfile", Value: userSastProfile},
		}}},
	)
	if result.Err() == mongo.ErrNoDocuments {
		return NoDocumentsErr
	}
	return result.Err()
}

func UpdateUserScore(userId *uuid.UUID, examId *string, userScoreMap *entity.UserScoreMap) error {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	fieldList := bson.D{}
	for problemId, userScore := range *userScoreMap {
		fieldList = append(fieldList, bson.E{
			Key:   fmt.Sprintf("scoreMap.%s.%s", *examId, problemId),
			Value: userScore,
		})
	}

	result := UserColl.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "userId", Value: userId}},
		bson.D{{Key: "$set", Value: fieldList}},
	)
	if result.Err() == mongo.ErrNoDocuments {
		return NoDocumentsErr
	}
	return result.Err()
}
