package response

// Response represents a standardized API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error details in the response
type ErrorInfo struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Success creates a success response
func JsonResponse(statusCode int, message string, data any, err error) Response {
	if err != nil {
		return Response{
			Success: false,
			Message: message,
			Data:    data,
			Error: &ErrorInfo{
				Type:    "error",
				Message: err.Error(),
			},
		}
	}
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}
