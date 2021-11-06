package response

import "github.com/TechnoHandOver/backend/internal/consts"

type Response struct {
	Code  consts.Code
	Data  interface{}
	Error error
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

func NewResponse(code consts.Code, data interface{}) *Response {
	return &Response{
		Code: code,
		Data: data,
	}
}

func NewEmptyResponse(code consts.Code) *Response {
	return &Response{
		Code: code,
	}
}

func NewErrorResponse(error_ error) *Response {
	return &Response{
		Code:  consts.InternalError,
		Error: error_,
	}
}
