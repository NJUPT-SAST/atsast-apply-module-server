package api

import (
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

func ReadConfig() *response.Response {
	config, err := service.ReadConfig()
	if err != nil {
		return response.Failed().SetMsg(err.Error())
	}
	return response.Success().SetData(config)
}
