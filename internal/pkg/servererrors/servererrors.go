package servererrors

import "errors"

var (
	ErrorInvalidServerResponseStatus = errors.New("invalid server response status")
	ErrorInternal                    = errors.New("internal server error")
	ErrorRecordNotFound              = errors.New("record not found")
)
