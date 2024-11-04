#!/bin/bash
echo "Script started running..."

# Enable debugging
set -x

# Basic test command to confirm AWS CLI functionality
aws --version

echo "AWS CLI is installed and working."



# #!/bin/bash
# echo "Script started running..."

# set -x
# set -e

# echo "Starting DynamoDB initialization..."

# # Check if required environment variables are set (optional if hardcoded)
# if [[ -z "$AWS_ACCESS_KEY_ID" || -z "$AWS_SECRET_ACCESS_KEY" || -z "$AWS_DEFAULT_REGION" ]]; then
#     echo "Please set AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, and AWS_DEFAULT_REGION environment variables."
#     exit 1
# fi

# # Configure AWS CLI before connecting to DynamoDB
# echo "Configuring AWS CLI..."

# # Create the .aws directory if it doesn't exist
# mkdir -p ~/.aws

# # Write the configuration files
# cat > ~/.aws/config <<EOL
# [default]
# region = $AWS_DEFAULT_REGION
# output = json
# EOL

# cat > ~/.aws/credentials <<EOL
# [default]
# aws_access_key_id = $AWS_ACCESS_KEY_ID
# aws_secret_access_key = $AWS_SECRET_ACCESS_KEY
# EOL

# echo "AWS CLI configured successfully."

# echo "AWS CLI configuration verified. Proceeding to create DynamoDB table..."

# # Create DynamoDB table
# echo "Creating DynamoDB table..."

aws dynamodb create-table \
    --table-name Images \
    --attribute-definitions \
        AttributeName=userID,AttributeType=S \
        AttributeName=imageID,AttributeType=S \
        AttributeName=uploadDate,AttributeType=S \
        AttributeName=analysisStatus,AttributeType=S \
        AttributeName=lastModified,AttributeType=S \
        AttributeName=dimensions,AttributeType=M \
        AttributeName=fileSize,AttributeType=N \
        AttributeName=fileType,AttributeType=S \
    --key-schema \
        AttributeName=userID,KeyType=HASH \
        AttributeName=imageID,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://dynamodb-local:8000

# echo "DynamoDB table created."
