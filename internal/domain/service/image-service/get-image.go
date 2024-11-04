package image_service

import (
	"context"
	"log"

	"github.com/ritik-patel05/image-player/internal/customerrors"
	"github.com/ritik-patel05/image-player/internal/domain/entity"
	image_db "github.com/ritik-patel05/image-player/internal/infrastructure/db/image-repo"
	"github.com/ritik-patel05/image-player/internal/models/dto"
)

func GetImage(ctx context.Context, reqDTO dto.GetImage) (*entity.ImageMetadata, error) {
	img, err := image_db.GetImageMetadata(ctx, reqDTO.ImageID)
	if err != nil {
		log.Printf("error getting image meta %v", err)
		return nil, err
	}

	if img.ImageID != reqDTO.ImageID {
		log.Printf("image not found %v", reqDTO.ImageID)
		return nil, customerrors.NewImageNotFoundError("image not found")
	}

	return img, nil
}
