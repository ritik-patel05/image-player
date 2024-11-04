package response

type CustomErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
