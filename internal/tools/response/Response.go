package response

type Response struct {
	Code       int
	JSONObject interface{}
}

func NewResponse(code int, jsonObject interface{}) *Response {
	return &Response{
		Code: code,
		JSONObject: jsonObject,
	}
}
