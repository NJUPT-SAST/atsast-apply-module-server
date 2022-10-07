package service

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/njupt-sast/atsast-apply-module-server/dao"
	"github.com/njupt-sast/atsast-apply-module-server/entity"
)

func ReadConfig() (*entity.Config, error) {
	config, err := dao.ReadConfig()
	if err == mongo.ErrNoDocuments {
		return nil, DocumentNotFoundError
	}

	if err != nil {
		return nil, err
	}
	return config, err
}
