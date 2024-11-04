package request

import "github.com/ritik-patel05/image-player/internal/customerrors"

type DeleteImage struct {
	ImageID string
}

func (req DeleteImage) Validate() error {
	if req.ImageID == "" {
		return customerrors.NewInvalidRequestError("imageId is required")
	}
	return nil
}
