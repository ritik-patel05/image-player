package dto

type UploadImageMetadata struct {
	UserID          string
	FileName        string
	DimensionWidth  int
	DimensionHeight int
	FileSize        int64
	FileType        string
}
