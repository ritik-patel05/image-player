package image_service

import (
	"context"
	"log"

	image_db "github.com/ritik-patel05/image-player/internal/infrastructure/db/image-repo"
	image_s3 "github.com/ritik-patel05/image-player/internal/infrastructure/s3/image-repo"
	"github.com/ritik-patel05/image-player/internal/models/dto"
)

func DeleteImage(ctx context.Context, reqDTO dto.DeleteImage) error {
	err := image_db.DeleteImageMetaData(ctx, reqDTO.ImageID)
	if err != nil {
		log.Printf("error deleting image meta from DB %v", err)
		return err
	}

	err = image_s3.DeleteImageMetaData(ctx, reqDTO.ImageID)
	if err != nil {
		log.Printf("error deleting image meta from S3 %v", err)
		return err
	}

	return nil
}
