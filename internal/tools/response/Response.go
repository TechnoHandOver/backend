package response

import "github.com/TechnoHandOver/backend/internal/consts"

type Response struct {
	Code  consts.Code
	Data  interface{}
	Error error
}

func NewResponse(code consts.Code, data interface{}) *Response {
	return &Response{
		Code: code,
		Data: data,
	}
}

func NewErrorResponse(code consts.Code, error_ error) *Response {
	return &Response{
		Code:  code,
		Error: error_,
	}
}

func NewEmptyResponse(code consts.Code) *Response {
	return &Response{
		Code: code,
	}
}
