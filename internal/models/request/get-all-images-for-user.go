package request

import "github.com/ritik-patel05/image-player/internal/customerrors"

type GetAllImagesForUser struct {
	UserID string
}

func (req GetAllImagesForUser) Validate() error {
	if req.UserID == "" {
		return customerrors.NewInvalidRequestError("userId is required")
	}
	return nil
}
