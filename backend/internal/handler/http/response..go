package handler

const (
	responseStatusSuccess = "success"
	responseStatusFail    = "fail"
	responseStatusError   = "error"
)

type JSONResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewResponse(status, message, err string) *JSONResponse {
	return &JSONResponse{
		Message: message,
		Status:  status,
		Error:   err,
	}
}
