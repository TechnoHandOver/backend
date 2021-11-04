package response

import "net/http"

type Response struct {
	Code  int
	Data  *DataResponse
	Error *error
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

func NewResponse(code int, jsonObject interface{}) *Response {
	return &Response{
		Code: code,
		Data: &DataResponse{
			Data: jsonObject,
		},
	}
}

func NewEmptyResponse(code int) *Response {
	return &Response{
		Code: code,
	}
}

func NewErrorResponse(error_ error) *Response {
	return &Response{
		Code:  http.StatusInternalServerError,
		Error: &error_,
	}
}
