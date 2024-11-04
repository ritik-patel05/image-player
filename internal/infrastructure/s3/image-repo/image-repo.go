package image_repo

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/bytedance/sonic"
	"github.com/ritik-patel05/image-player/internal/constants"
	"github.com/ritik-patel05/image-player/internal/domain/entity"
	"github.com/ritik-patel05/image-player/internal/driver"
)

func UploadImageMetadata(ctx context.Context, imageMetaData entity.ImageMetadata) error {
	s3Client := driver.GetS3Client()

	jsonImageData, err := sonic.Marshal(imageMetaData)
	if err != nil {
		log.Printf("UploadImageMetadata: Error Marshalling image metadata")
		return err
	}

	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(constants.IMAGES_BUCKET),
		Key:         aws.String(imageMetaData.ImageID),
		Body:        bytes.NewReader(jsonImageData),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		log.Printf("UploadImageMetadata: Error Upload image metadata to s3")
		return err
	}

	return nil
}

func DeleteImageMetaData(ctx context.Context, imageID string) error {
	s3Client := driver.GetS3Client()

	_, err := s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(constants.IMAGES_BUCKET),
		Key:    aws.String(imageID),
	})
	if err != nil {
		log.Printf("DeleteImageMetaData: Error Deleting image metadata from s3")
		return err
	}

	return nil
}
