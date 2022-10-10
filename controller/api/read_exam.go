package api

import (
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

func ReadExamList() *response.Response {
	examList, err := service.ReadExamList()
	if err != nil {
		return response.Failed()
	}

	return response.Success().SetData(examList)
}
