#!/bin/bash
echo "########### Setting region as env variable ##########"
export AWS_REGION=sa-east-1

echo "########### Setting up localstack profile ###########"
aws configure set aws_access_key_id access_key --profile=localstack
aws configure set aws_secret_access_key secret_key --profile=localstack
aws configure set region $AWS_REGION --profile=localstack

echo "########### Setting default profile ###########"
export AWS_DEFAULT_PROFILE=localstack

echo "########### Setting SQS and S3 names as env variables ###########"
export UPLOAD_IMAGES_EVENT_SQS=upload-images-event-sqs
export DELETE_IMAGES_EVENT_SQS=delete-images-event-sqs
export BUCKET_NAME=images

echo "########### Creating upload file event SQS ###########"
aws --endpoint-url=http://localstack:4566 sqs create-queue --queue-name $UPLOAD_IMAGES_EVENT_SQS

echo "########### ARN for upload file event SQS ###########"
UPLOAD_IMAGES_EVENT_SQS_ARN=$(aws --endpoint-url=http://localstack:4566 sqs get-queue-attributes\
                  --attribute-name QueueArn --queue-url=http://localhost:4566/000000000000/"$UPLOAD_IMAGES_EVENT_SQS"\
                  |  sed 's/"QueueArn"/\n"QueueArn"/g' | grep '"QueueArn"' | awk -F '"QueueArn":' '{print $2}' | tr -d '"' | xargs)

echo "########### Creating delete file event SQS ###########"
aws --endpoint-url=http://localstack:4566 sqs create-queue --queue-name $DELETE_IMAGES_EVENT_SQS

echo "########### ARN for delete file event SQS ###########"
DELETE_IMAGES_EVENT_SQS_ARN=$(aws --endpoint-url=http://localstack:4566 sqs get-queue-attributes\
                  --attribute-name QueueArn --queue-url=http://localhost:4566/000000000000/"$DELETE_IMAGES_EVENT_SQS"\
                  |  sed 's/"QueueArn"/\n"QueueArn"/g' | grep '"QueueArn"' | awk -F '"QueueArn":' '{print $2}' | tr -d '"' | xargs)

echo "########### Listing queues ###########"
aws --endpoint-url=http://localhost:4566 sqs list-queues

echo "########### Create S3 bucket ###########"
aws --endpoint-url=http://localhost:4566 s3api create-bucket\
    --bucket $BUCKET_NAME --region $AWS_REGION\
    --create-bucket-configuration LocationConstraint=$AWS_REGION

echo "########### List S3 bucket ###########"
aws --endpoint-url=http://localhost:4566 s3api list-buckets

echo "########### Set S3 bucket notification configurations ###########"
aws --endpoint-url=http://localhost:4566 s3api put-bucket-notification-configuration\
    --bucket $BUCKET_NAME\
    --notification-configuration  '{
                                      "QueueConfigurations": [
                                         {
                                           "QueueArn": "'"$UPLOAD_IMAGES_EVENT_SQS_ARN"'",
                                           "Events": ["s3:ObjectCreated:*"]
                                         },
                                         {
                                            "QueueArn": "'"$DELETE_IMAGES_EVENT_SQS_ARN"'",
                                            "Events": ["s3:ObjectRemoved:*"]
                                          }
                                       ]
                                     }'

echo "########### Get S3 bucket notification configurations ###########"
aws --endpoint-url=http://localhost:4566 s3api get-bucket-notification-configuration\
    --bucket $BUCKET_NAME

echo "########### Create DynamoDB table with GSI ###########"
awslocal dynamodb create-table \
    --table-name Images \
    --key-schema AttributeName=ImageID,KeyType=HASH \
    --attribute-definitions AttributeName=ImageID,AttributeType=S AttributeName=UserID,AttributeType=S \
    --global-secondary-indexes '[
        {
            "IndexName": "UserIDIndex",
            "KeySchema": [
                {
                    "AttributeName": "UserID",
                    "KeyType": "HASH"
                }
            ],
            "Projection": {
                "ProjectionType": "ALL"
            },
            "ProvisionedThroughput": {
                "ReadCapacityUnits": 5,
                "WriteCapacityUnits": 5
            }
        }
    ]' \
    --table-class STANDARD \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5



echo "########### List DynamoDB tables ###########"
aws --endpoint-url=http://localhost:4566 dynamodb list-tables
