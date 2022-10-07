package stdrsp

const (
	ErrMsgSuffix           = ", pls try again later or contact the admin"
	AuthErrMsg             = "auth error"
	BadRequestErrMsg       = "bad request error"
	PermissionDeniedErrMsg = "permission denied error"
	CallDatabaseErrMsg     = "calling database error" + ErrMsgSuffix
	CallWechatApiErrMsg    = "calling wechat api error" + ErrMsgSuffix
	NotFoundErrMsg         = "something not found" + ErrMsgSuffix
)
