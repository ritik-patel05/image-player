package image_service

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ritik-patel05/image-player/internal/domain/entity"
	image_db "github.com/ritik-patel05/image-player/internal/infrastructure/db/image-repo"
	image_s3 "github.com/ritik-patel05/image-player/internal/infrastructure/s3/image-repo"
	"github.com/ritik-patel05/image-player/internal/models/dto"
	"github.com/ritik-patel05/image-player/internal/models/response"
)

func UploadImageMetadata(ctx context.Context, reqDTO dto.UploadImageMetadata) (*response.UploadImageMetadata, error) {
	currTime := time.Now()
	analysisStatus := "UPLOADED"

	imageID := uuid.NewString()

	imageMetadata := entity.NewImageMetadata(imageID, reqDTO.UserID, &reqDTO.FileName,
		&currTime, &currTime, &reqDTO.DimensionWidth, &reqDTO.DimensionHeight,
		&reqDTO.FileSize, &reqDTO.FileType, &analysisStatus, nil)

	err := image_db.InsertImageMetadata(ctx, *imageMetadata)
	if err != nil {
		log.Printf("error inserting image meta to DB %v", err)
		return nil, err
	}

	err = image_s3.UploadImageMetadata(ctx, *imageMetadata)
	if err != nil {
		log.Printf("error inserting image meta to S3 %v", err)
		return nil, err
	}

	imageUploadResp := response.UploadImageMetadata{
		ImageID: imageID,
	}

	return &imageUploadResp, nil
}
