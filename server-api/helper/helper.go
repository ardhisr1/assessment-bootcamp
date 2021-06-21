package helper

type ApiResponse struct {
	Code    int
	Message string
	Data    interface{}
}

type FailureResponse struct {
	Code    int
	Message string
	Error   interface{}
}

func APIResponse(code int, msg string, data interface{}) ApiResponse {
	var response = ApiResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	return response
}

func FailResponse(code int, msg string, Error interface{}) FailureResponse {
	var response = FailureResponse{
		Code:    code,
		Message: msg,
		Error:   Error,
	}
	return response
}
