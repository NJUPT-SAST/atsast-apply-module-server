package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/njupt-sast/atsast-apply-module-server/entity"
)

func ReadConfig() (*entity.Config, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	config := entity.Config{}
	if err := ConfigColl.FindOne(ctx, bson.D{}).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
