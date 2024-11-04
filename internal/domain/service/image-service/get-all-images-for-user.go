package image_service

import (
	"context"
	"log"

	image_db "github.com/ritik-patel05/image-player/internal/infrastructure/db/image-repo"
	"github.com/ritik-patel05/image-player/internal/models/dto"
	"github.com/ritik-patel05/image-player/internal/models/response"
)

func GetAllImagesForUser(ctx context.Context, reqDTO dto.GetAllImagesForUser) (*response.GetAllImagesForUser, error) {
	images, err := image_db.FindAllImagesForUser(ctx, reqDTO.UserID)
	if err != nil {
		log.Printf("error finding all images for a user %s - %v", reqDTO.UserID, err)
		return nil, err
	}

	return &response.GetAllImagesForUser{
		Images: images,
	}, nil
}
