package image_service

import (
	"context"
	"log"

	"github.com/ritik-patel05/image-player/internal/domain/entity"
	image_db "github.com/ritik-patel05/image-player/internal/infrastructure/db/image-repo"
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

	return nil
}
