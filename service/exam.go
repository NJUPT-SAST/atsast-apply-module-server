package service

import (
	"github.com/njupt-sast/atsast-apply-module-server/model/dao"
	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
)

func ReadExamList() ([]entity.Exam, error) {
	examList, err := dao.ReadExamList()
	if err == dao.NoDocumentsErr {
		return examList, nil
	}
	if err != nil {
		return nil, CallDatabaseErr
	}
	return examList, nil
}
