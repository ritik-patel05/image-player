package driver

import (
	"context"
	"log"

	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ritik-patel05/image-player/config"
	"github.com/ritik-patel05/image-player/internal/constants"
)

var s3Client *s3.Client

func init() {
	s3Client = newS3Client()
}

func newS3Client() *s3.Client {
	// Credentials - Reads env variables (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN)
	cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion(constants.AWS_REGION))
	if err != nil {
		log.Fatalf("Failed to load S3 configuration: %v", err)
	}

	if config.GetConfig().ActiveEnv == constants.DEV {
		localStackEndpoint := constants.AWS_LOCALSTACK_ENDPOINT

		return s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
			o.BaseEndpoint = &localStackEndpoint
		})
	}

	return s3.NewFromConfig(cfg)
}

func GetS3Client() *s3.Client {
	return s3Client
}
