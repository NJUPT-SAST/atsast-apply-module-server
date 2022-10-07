package controller

import (
	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

func readExamList() *stdrsp.Result {
	examList, err := service.ReadExamList()
	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
	}

	return stdrsp.Response(stdrsp.Success).Data(examList)
}
