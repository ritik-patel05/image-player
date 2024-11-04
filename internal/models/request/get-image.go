package request

import "github.com/ritik-patel05/image-player/internal/customerrors"

type GetImage struct {
	ImageID string
}

func (req GetImage) Validate() error {
	if req.ImageID == "" {
		return customerrors.NewInvalidRequestError("imageId is required")
	}
	return nil
}
