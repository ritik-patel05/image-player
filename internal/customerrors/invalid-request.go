package customerrors

import "net/http"

const InvalidRequestErrorCode = "INVALID_REQUEST"

func NewInvalidRequestError(message string) error {
	return &CustomError{
		Message:        message,
		HttpStatusCode: http.StatusBadRequest,
		ErrorCode:      InvalidRequestErrorCode,
	}
}
