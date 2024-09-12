package middleware

type Response struct {
	Success    bool   `json:"success"`
	ErrCode    string `json:"errCode,omitempty"`
	ErrMessage string `json:"errMessage,omitempty"`
}

// NewSuccessResponse 构建一个成功响应
func NewSuccessResponse() *Response {
	return &Response{
		Success: true,
	}
}

// NewErrorResponse 构建一个失败响应
func NewErrorResponse(errCode, errMessage string) *Response {
	return &Response{
		Success:    false,
		ErrCode:    errCode,
		ErrMessage: errMessage,
	}
}

// SingleResponse 是一个包含单个数据项的响应结构体
type SingleResponse struct {
	Response
	Data any `json:"data,omitempty"`
}

// NewSingleResponse 构建一个包含单个数据项的成功响应
func NewSingleResponse(data any) *SingleResponse {
	return &SingleResponse{
		Response: Response{
			Success: true,
		},
		Data: data,
	}
}

// MultiResponse 是一个包含多个数据项的响应结构体
type MultiResponse struct {
	Response
	Data []interface{} `json:"data,omitempty"`
}

// NewMultiResponse 构建一个包含多个数据项的成功响应
func NewMultiResponse(data []any) *MultiResponse {
	return &MultiResponse{
		Response: Response{
			Success: true,
		},
		Data: data,
	}
}
