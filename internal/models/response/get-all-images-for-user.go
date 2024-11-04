package response

import "github.com/ritik-patel05/image-player/internal/domain/entity"

type GetAllImagesForUser struct {
	Images []entity.ImageMetadata `json:"images"`
}
