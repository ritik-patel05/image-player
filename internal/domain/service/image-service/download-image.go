package image_service

import (
	"context"

	"github.com/ritik-patel05/image-player/internal/models/dto"
)

func DownloadImage(ctx context.Context, reqDTO dto.DownloadImage) ([]byte, error) {
	// Not implemented because currently we are not uploading file to s3.
	// Its just a metadata being uploaded to s3. You can use GetImage API to fetch image Image Meta.
	return nil, nil
}
