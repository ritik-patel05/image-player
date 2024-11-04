package customerrors

const PermissionDeniedErrorCode = "PERMISSION_DENIED"

func NewPermissionDenied(message string) *CustomError {
	return &CustomError{
		Message:        message,
		HttpStatusCode: 403,
		ErrorCode:      PermissionDeniedErrorCode,
	}
}
