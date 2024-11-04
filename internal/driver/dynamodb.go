package driver

import (
	"context"
	"log"

	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/ritik-patel05/image-player/config"
	"github.com/ritik-patel05/image-player/internal/constants"
)

var dynamoDBClient *dynamodb.Client

func init() {
	dynamoDBClient = newDynamoDBClient()
}

func newDynamoDBClient() *dynamodb.Client {
	// Credentials - Reads env variables (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN)
	cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion(constants.AWS_REGION))
	if err != nil {
		log.Fatalf("Failed to load dynamoDB configuration: %v", err)
	}

	if config.GetConfig().ActiveEnv == constants.DEV {
		localStackEndpoint := constants.AWS_LOCALSTACK_ENDPOINT

		return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = &localStackEndpoint
		})
	}

	return dynamodb.NewFromConfig(cfg)
}

func GetDynamoDBClient() *dynamodb.Client {
	return dynamoDBClient
}
