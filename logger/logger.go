package logger

import (
	"os"

	"github.com/google/uuid"
)

var defaultRequestLogger RequestLogger

func init() {
	if err := os.MkdirAll("logs", os.ModeAppend); err != nil {
		panic(err)
	}

	logfile, err := os.OpenFile("logs/request.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defaultRequestLogger = newRequestLogger()
	defaultRequestLogger.SetOutput(logfile)

}

func LogRequest(api string, userId *uuid.UUID, request interface{}) {
	defaultRequestLogger.logRequest(api, userId, request)
}
