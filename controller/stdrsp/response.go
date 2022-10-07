package stdrsp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusCode int

const (
	Success    StatusCode = 200
	Failed     StatusCode = 201
	BadRequest StatusCode = 400
	AuthFailed StatusCode = 401
)

type Result struct {
	ResultCode StatusCode  `json:"code"`
	ResultData interface{} `json:"data"`
	ResultMsg  string      `json:"message"`
}

func Response(code StatusCode) *Result {
	return &Result{
		ResultCode: code,
	}
}

func (result *Result) Data(data interface{}) *Result {
	result.ResultData = data
	return result
}

func (result *Result) Msg(msg string) *Result {
	result.ResultMsg = msg
	return result
}

func ResponseResult(c *gin.Context, result *Result) {
	c.JSON(http.StatusOK, result)
}
