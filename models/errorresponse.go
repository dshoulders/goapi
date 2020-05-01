package models

type ErrorResponse struct {
	Status  `json:"status"`
	Message string `json:"message"`
}

func CreateErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Error, message}
}
