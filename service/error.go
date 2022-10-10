package service

import "errors"

var (
	CallDatabaseErr     = errors.New("calling database error")
	NotFoundErr         = errors.New("not found error")
	PermissionDeniedErr = errors.New("permission denied")
)
