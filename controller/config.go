package controller

import (
	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

func readConfig() *stdrsp.Result {
	config, err := service.ReadConfig()
	if err == service.DocumentNotFoundError {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
	}

	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
	}

	return stdrsp.Response(stdrsp.Success).Data(config)
}
