package middleware

import (
	"fx-web/internal/consts"
	"strconv"
)

type Response struct {
	Success    bool   `json:"success"`
	ErrCode    string `json:"errCode,omitempty"`
	ErrMessage string `json:"errMessage,omitempty"`
	Type       string `json:"type,omitempty"`
}

// NewSuccessResponse 构建一个成功响应
func NewSuccessResponse() *Response {
	return &Response{
		Success: true,
		Type:    "BaseResponse",
	}
}

// NewErrorResponse 构建一个失败响应
func NewErrorResponse(errCode, errMessage string) *Response {
	return &Response{
		Success:    false,
		ErrCode:    errCode,
		ErrMessage: errMessage,
		Type:       "ErrorResponse",
	}
}

type CommonResponse struct {
	Response
	Data any `json:"data,omitempty"`
}

// NewCommonResponse 构建一个包含单个数据项的成功响应
func NewCommonResponse(data any) *CommonResponse {
	return &CommonResponse{
		Response: Response{
			Success: true,
			Type:    "CommonResponse",
		},
		Data: data,
	}
}

type CommonErrorResponse struct {
	Response
	Data any `json:"data,omitempty"`
}

func NewCommonErrorResponse(error consts.CommonError) *CommonErrorResponse {
	return &CommonErrorResponse{
		Response: Response{
			Success:    false,
			ErrCode:    strconv.Itoa(error.Code),
			ErrMessage: error.Msg,
			Type:       "ErrorResponse",
		},
		Data: error.Data,
	}
}

func NewCommonErrorResponseWithData(error consts.CommonError, data any) *CommonErrorResponse {
	return &CommonErrorResponse{
		Response: Response{
			Success:    false,
			ErrCode:    strconv.Itoa(error.Code),
			ErrMessage: error.Msg,
			Type:       "ErrorResponse",
		},
		Data: data,
	}
}
