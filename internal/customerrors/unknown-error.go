package customerrors

import "net/http"

const unknownErrorCode = "INTERNAL_SERVER_ERROR"

func NewUnknownError() error {
	return &CustomError{
		Message:        "internal server error",
		HttpStatusCode: http.StatusInternalServerError,
		ErrorCode:      unknownErrorCode,
	}
}
