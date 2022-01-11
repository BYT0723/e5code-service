package errorx

import "fmt"

const (
	SQLError = 100

	ServiceError = 200

	APIError = 300

	// API ERROR
	EncryptPwdError = 301
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("%v: %s", e.Code, e.Msg)
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func (e *CodeError) Belong(parent *CodeError) bool {
	return e.Code-parent.Code > 0 && e.Code-parent.Code < 100
}
