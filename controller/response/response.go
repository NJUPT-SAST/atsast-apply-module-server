package response

type StatusCode int

var (
	StatusSuccess StatusCode = 200
	StatusFailed  StatusCode = 500
)

type Response struct {
	Code StatusCode `json:"code"`
	Data any		`json:"data"`
	Msg  *string	`json:"message"`
}

func Success() *Response {
	return &Response{Code: StatusSuccess}
}

func Failed() *Response {
	return &Response{Code: StatusFailed}
}

func (response *Response) SetData(data any) *Response {
	response.Data = data
	return response
}

func (response *Response) SetMsg(msg string) *Response {
	response.Msg = &msg
	return response
}
