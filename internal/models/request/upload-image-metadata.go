package request

import "github.com/ritik-patel05/image-player/internal/customerrors"

type UploadImageMetadata struct {
	ImageID         string `json:"imageId"`
	UserID          string `json:"userId"`
	FileName        string `json:"fileName"`
	DimensionWidth  int    `json:"dimensionWidth"`
	DimensionHeight int    `json:"dimensionHeight"`
	FileSize        int64  `json:"fileSize"`
	FileType        string `json:"fileType"`
}

func (req UploadImageMetadata) Validate() error {
	if req.UserID == "" {
		return customerrors.NewInvalidRequestError("userId is required")
	}
	return nil
}
