package errors

type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e BizError) Error() string {
	return e.Message
}
func NewBizError(code int, msg string) *BizError {
	return &BizError{
		Code:    code,
		Message: msg,
	}
}
