package service

import (
	"github.com/njupt-sast/atsast-apply-module-server/model/dao"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
)

func ReadConfig() (*entity.Config, error) {
	config, err := dao.ReadConfig()
	if err == dao.NoDocumentsErr {
		return nil, NotFoundErr
	}
	if err != nil {
		return nil, CallDatabaseErr
	}
	return config, nil
}
