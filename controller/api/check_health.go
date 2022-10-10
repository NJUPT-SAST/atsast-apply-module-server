package api

import (
	"github.com/njupt-sast/atsast-apply-module-server/controller/response"
)

func CheckHealth() *response.Response {
	return response.Success().SetMsg("success")
}
