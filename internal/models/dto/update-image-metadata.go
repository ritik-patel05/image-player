package dto

type UpdateImageMetadata struct {
	ImageID         string
	UserID          string
	FileName        *string
	DimensionWidth  *int
	DimensionHeight *int
	FileSize        *int64
	FileType        *string
	AnalysisStatus  *string
	S3URL           *string
}
