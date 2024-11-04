package entity

import (
	"time"
)

type ImageMetadata struct {
	ImageID         string     `dynamodbav:"ImageID"         json:"imageId"`
	UserID          string     `dynamodbav:"UserID"          json:"userId"`
	FileName        *string    `dynamodbav:"FileName"        json:"fileName"`
	UploadDate      *time.Time `dynamodbav:"UploadDate"      json:"uploadDate"`
	LastUpdatedAt   *time.Time `dynamodbav:"LastUpdatedAt"   json:"lastUpdatedAt"`
	DimensionWidth  *int       `dynamodbav:"DimensionWidth"  json:"dimensionWidth"`
	DimensionHeight *int       `dynamodbav:"DimensionHeight" json:"dimensionHeight"`
	FileSize        *int64     `dynamodbav:"FileSize"        json:"fileSize"`
	FileType        *string    `dynamodbav:"FileType"        json:"fileType"`
	AnalysisStatus  *string    `dynamodbav:"AnalysisStatus"  json:"analysisStatus"`
	S3URL           *string    `dynamodbav:"S3URL"           json:"s3Url"`
}

func NewImageMetadata(imageID string, userID string, fileName *string, uploadDate *time.Time,
	lastUpdatedAt *time.Time, dimensionWidth *int, dimensionHeight *int, fileSize *int64,
	fileType *string, analysisStatus *string, s3URL *string) *ImageMetadata {

	imageMetaData := ImageMetadata{
		ImageID:         imageID,
		UserID:          userID,
		FileName:        fileName,
		UploadDate:      uploadDate,
		LastUpdatedAt:   lastUpdatedAt,
		DimensionWidth:  dimensionWidth,
		DimensionHeight: dimensionHeight,
		FileSize:        fileSize,
		FileType:        fileType,
		AnalysisStatus:  analysisStatus,
		S3URL:           s3URL,
	}

	return &imageMetaData
}
