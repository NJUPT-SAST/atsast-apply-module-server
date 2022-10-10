package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
)

func ReadConfig() (*entity.Config, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	config := entity.Config{}
	err := ConfigColl.FindOne(ctx, bson.D{}).Decode(&config)
	if err == mongo.ErrNoDocuments {
		return nil, NoDocumentsErr
	}
	if err != nil {
		return nil, err
	}
	return &config, nil
}
