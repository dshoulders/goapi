package models

type SuccessResponse struct {
	Status `json:"status"`
	Data   JSON `json:"data"`
}

func CreateSuccessResponse(response JSON) SuccessResponse {
	return SuccessResponse{Success, response}
}
