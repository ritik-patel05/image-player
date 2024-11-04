package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ritik-patel05/image-player/internal/customerrors"
	"github.com/ritik-patel05/image-player/internal/models/response"
)

func PrepareErrorResponse(c *gin.Context, err error) {
	var customError *customerrors.CustomError
	if errors.As(err, &customError) {
		c.JSON(customError.HttpStatusCode, response.CustomErrorResponse{
			Message:   customError.Error(),
			ErrorCode: customError.ErrorCode,
		})
		return
	}

	var unknownError *customerrors.CustomError
	_ = errors.As(customerrors.NewUnknownError(), &unknownError)
	c.JSON(unknownError.HttpStatusCode, response.CustomErrorResponse{
		Message:   unknownError.Error(),
		ErrorCode: unknownError.ErrorCode,
	})

	return
}
