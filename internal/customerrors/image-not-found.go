package customerrors

import "net/http"

const ImageNotFoundErrorCode = "IMAGE_NOT_FOUND"

func NewImageNotFoundError(message string) error {
	return &CustomError{
		Message:        message,
		HttpStatusCode: http.StatusNotFound,
		ErrorCode:      ImageNotFoundErrorCode,
	}
}
