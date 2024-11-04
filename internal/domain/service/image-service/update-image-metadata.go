package image_service

import (
	"context"
	"log"

	"github.com/ritik-patel05/image-player/internal/domain/entity"
	image_db "github.com/ritik-patel05/image-player/internal/infrastructure/db/image-repo"
	image_s3 "github.com/ritik-patel05/image-player/internal/infrastructure/s3/image-repo"
	"github.com/ritik-patel05/image-player/internal/models/dto"
)

func UpdateImageMetadata(ctx context.Context, reqDTO dto.UpdateImageMetadata) error {
	imgMeta := entity.ImageMetadata{
		ImageID:         reqDTO.ImageID,
		FileName:        reqDTO.FileName,
		DimensionWidth:  reqDTO.DimensionWidth,
		DimensionHeight: reqDTO.DimensionHeight,
		FileSize:        reqDTO.FileSize,
		FileType:        reqDTO.FileType,
		AnalysisStatus:  reqDTO.AnalysisStatus,
		S3URL:           reqDTO.S3URL,
	}

	err := image_db.UpdateImageMetadata(ctx, imgMeta)
	if err != nil {
		log.Printf("error updating image meta %v", err)
		return err
	}

	newImgMeta, err := image_db.GetImageMetadata(ctx, imgMeta.ImageID)
	if err != nil {
		log.Printf("error while getting new image meta %v", err)
		return err
	}

	err = image_s3.DeleteImageMetaData(ctx, imgMeta.ImageID)
	if err != nil {
		log.Printf("error deleting image meta from S3 %v", err)
		return err
	}

	err = image_s3.UploadImageMetadata(ctx, *newImgMeta)
	if err != nil {
		log.Printf("error inserting image meta to S3 %v", err)
		return err
	}

	return nil
}
