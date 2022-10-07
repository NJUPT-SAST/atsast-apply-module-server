package logger

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RequestLogger struct {
	*logrus.Logger
}

type RequestLoggerFormatter struct {
	TimestampFormat string
}

func newRequestLogger() RequestLogger {
	logger := RequestLogger{logrus.New()}
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&RequestLoggerFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func (f *RequestLoggerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	time := entry.Time.Format(f.TimestampFormat)
	level := strings.ToUpper(entry.Level.String())
	return []byte(fmt.Sprintf("[%s] %s | %s\n", time, level, entry.Message)), nil
}

func (logger *RequestLogger) logRequest(api string, userId *uuid.UUID, request interface{}) {
	userIdString := "null"
	if userId != nil {
		userIdString = userId.String()
	}

	requestString := "null"
	if request != nil {
		requestByte, _ := json.Marshal(request)
		requestString = string(requestByte)
	}

	logger.Info(fmt.Sprintf("%s: %s, %s", api, userIdString, requestString))
}
