package service

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/njupt-sast/atsast-apply-module-server/dao"
	"github.com/njupt-sast/atsast-apply-module-server/entity"
)

func ReadExamList() ([]entity.Exam, error) {
	examList, err := dao.ReadExamList()
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return examList, nil
}
