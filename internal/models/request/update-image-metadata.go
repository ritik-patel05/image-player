package request

import "github.com/ritik-patel05/image-player/internal/customerrors"

type UpdateImageMetadata struct {
	ImageID         string
	UserID          string
	FileName        *string `json:"fileName"`
	DimensionWidth  *int    `json:"dimensionWidth"`
	DimensionHeight *int    `json:"dimensionHeight"`
	FileSize        *int64  `json:"fileSize"`
	FileType        *string `json:"fileType"`
	AnalysisStatus  *string `json:"analysisStatus"`
	S3URL           *string `json:"s3Url"`
}

func (req UpdateImageMetadata) Validate() error {
	if req.ImageID == "" {
		return customerrors.NewInvalidRequestError("imageId is required")
	}
	if req.UserID == "" {
		return customerrors.NewInvalidRequestError("userId is required")
	}
	return nil
}
