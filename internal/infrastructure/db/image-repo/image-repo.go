package image_repo

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ritik-patel05/image-player/internal/constants"
	"github.com/ritik-patel05/image-player/internal/domain/entity"
	"github.com/ritik-patel05/image-player/internal/driver"
)

func InsertImageMetadata(ctx context.Context, imageMetaData entity.ImageMetadata) error {
	dynamoDBClient := driver.GetDynamoDBClient()

	data, err := attributevalue.MarshalMap(imageMetaData)
	if err != nil {
		return err
	}

	_, err = dynamoDBClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(constants.IMAGES_TABLE),
		Item:      data,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateImageMetadata(ctx context.Context, imgMeta entity.ImageMetadata) error {
	dynamoDBClient := driver.GetDynamoDBClient()

	currTime := time.Now()
	update := expression.Set(expression.Name("LastUpdatedAt"), expression.Value(currTime))

	if imgMeta.FileName != nil {
		update.Set(expression.Name("FileName"), expression.Value(imgMeta.FileName))
	}
	if imgMeta.DimensionWidth != nil {
		update.Set(expression.Name("DimensionWidth"), expression.Value(imgMeta.DimensionWidth))
	}
	if imgMeta.DimensionHeight != nil {
		update.Set(expression.Name("DimensionHeight"), expression.Value(imgMeta.DimensionHeight))
	}
	if imgMeta.FileSize != nil {
		update.Set(expression.Name("FileSize"), expression.Value(imgMeta.FileSize))
	}
	if imgMeta.FileType != nil {
		update.Set(expression.Name("FileType"), expression.Value(imgMeta.FileType))
	}
	if imgMeta.AnalysisStatus != nil {
		update.Set(expression.Name("AnalysisStatus"), expression.Value(imgMeta.AnalysisStatus))
	}
	if imgMeta.S3URL != nil {
		update.Set(expression.Name("S3URL"), expression.Value(imgMeta.S3URL))
	}

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return err
	}

	pk, err := getPrimaryKey(imgMeta.ImageID)
	if err != nil {
		return err
	}

	_, err = dynamoDBClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(constants.IMAGES_TABLE),
		Key:                       pk,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	})
	if err != nil {
		return err
	}

	return err
}

func DeleteImageMetaData(ctx context.Context, imageID string) error {
	dynamoDBClient := driver.GetDynamoDBClient()

	pk, err := getPrimaryKey(imageID)
	if err != nil {
		return err
	}

	_, err = dynamoDBClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(constants.IMAGES_TABLE),
		Key:       pk,
	})
	if err != nil {
		return err
	}

	return err
}

func GetImageMetadata(ctx context.Context, imageID string) (*entity.ImageMetadata, error) {
	dynamoDBClient := driver.GetDynamoDBClient()

	pk, err := getPrimaryKey(imageID)
	if err != nil {
		return nil, err
	}

	result, err := dynamoDBClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(constants.IMAGES_TABLE),
		Key:       pk,
	})
	if err != nil {
		return nil, err
	}

	imageMetaData := entity.ImageMetadata{}
	err = attributevalue.UnmarshalMap(result.Item, &imageMetaData)
	if err != nil {
		return nil, err
	}

	return &imageMetaData, nil
}

func FindAllImagesForUser(ctx context.Context, userID string) ([]entity.ImageMetadata, error) {
	dynamoDBClient := driver.GetDynamoDBClient()

	keyEx := expression.Key("UserID").Equal(expression.Value(userID))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}

	queryPaginator := dynamodb.NewQueryPaginator(dynamoDBClient, &dynamodb.QueryInput{
		TableName:                 aws.String(constants.IMAGES_TABLE),
		IndexName:                 aws.String(constants.USER_ID_GSI_INDEX),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	imagesMetaData := []entity.ImageMetadata{}
	for queryPaginator.HasMorePages() {
		response, err := queryPaginator.NextPage(ctx)
		if err != nil {
			log.Printf("Failed to get query response. Here's why: %v\n", err)
			break
		}

		imagesMetaDataPage := []entity.ImageMetadata{}
		err = attributevalue.UnmarshalListOfMaps(response.Items, &imagesMetaDataPage)
		if err != nil {
			log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
			break
		}
		imagesMetaData = append(imagesMetaData, imagesMetaDataPage...)
	}

	return imagesMetaData, nil
}

func getPrimaryKey(imageID string) (map[string]types.AttributeValue, error) {
	iID, err := attributevalue.Marshal(imageID)
	if err != nil {
		return nil, err
	}

	return map[string]types.AttributeValue{
		"ImageID": iID,
	}, nil
}
