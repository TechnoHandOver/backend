package response

import "github.com/TechnoHandOver/backend/internal/models"

type Response struct {
	Code       int
	JSONObject interface{}
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error *models.Error `json:"error"`
}

func NewResponse(code int, jsonObject interface{}) *Response {
	return &Response{
		Code: code,
		JSONObject: DataResponse{
			Data: jsonObject,
		},
	}
}

func NewErrorResponse(code int, error *models.Error) *Response {
	return &Response{
		Code: code,
		JSONObject: ErrorResponse{
			Error: error,
		},
	}
}
