package request

import "github.com/ritik-patel05/image-player/internal/customerrors"

type DownloadImage struct {
	ImageID string
}

func (req DownloadImage) Validate() error {
	if req.ImageID == "" {
		return customerrors.NewInvalidRequestError("imageId is required")
	}
	return nil
}
