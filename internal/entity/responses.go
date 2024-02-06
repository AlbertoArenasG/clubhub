package entity

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessListResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func NewSuccessListResponse(message string, data interface{}, meta interface{}) *SuccessListResponse {
	return &SuccessListResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func NewErrorResponse(err string) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Error:   err,
	}
}
