package response

type Response struct {
	Code       int
	JSONObject WrappedResponse
}

type WrappedResponse struct {
	Data interface{} `json:"data"`
}

func NewResponse(code int, jsonObject interface{}) *Response {
	return &Response{
		Code: code,
		JSONObject: WrappedResponse{
			Data: jsonObject,
		},
	}
}
