# Image-Player

## How to start the service Locally
1. [Run the docker compose command by following steps mentioned here](docs/development/setup-aws-locally.md#prerequisites)
2. Start the golang service using the following commands in root directory.
    ```bash
    ACTIVE_ENV=DEV AWS_SECRET_ACCESS_KEY=secret_key AWS_ACCESS_KEY_ID=access_key go run ./cmd/image-service
    ```


A platform to manage images with features for uploading, retrieving, updating, and deleting images. This API provides endpoints for handling image metadata and integrating with AWS S3 for secure storage.

## API Documentation

**Postman Collection** - https://drive.google.com/file/d/14R_c55S5c4g4hRACAsJefw-B_w1vVBjN/view?usp=drive_link 

Here's a detailed contract for each API endpoint:

---

### 1. **Upload Image Metadata**

- **Endpoint**: `POST /images/upload`
- **Method**: `POST`
- **Description**: Accepts image metadata and  in future if we want to support file uploads from client - in response we can generate a pre-signed S3 URL for uploading the image.
- **Request Body**:
  ```json
  {
    "fileName": "string",
    "dimensionWidth": "int",
    "dimensionHeight": "int",
    "fileSize": "int64",
    "fileType": "string"
  }
  ```
- **Response**:
  - **200 OK**:
    ```json
    {
      "imageID": "strng",
      "uploadUrl": "string" // For Future, Pre-signed S3 URL for image upload
    }
    ```
  - **400 Bad Request**: If the request body is invalid.
    ```json
    {
      "error": "invalid request body"
    }
    ```

---

### 2. **List All Images for a User**

- **Endpoint**: `GET /images/user/{userID}`
- **Method**: `GET`
- **Description**: Retrieves a paginated list of images for a specific user.
- **Path Parameters**:
  - `userID` (string) – The ID of the user.
- **Response**:
  - **200 OK**:
    ```json
    {
      "images": [
        {
          "imageId": "string",
          "fileName": "string",
          "uploadDate": "string",
          "fileSize": "int64",
          "fileType": "string"
        }
      ],
      "totalCount": "int"
    }
    ```
  - **404 Not Found**: If no images are found for the user.
    ```json
    {
      "error": "No images found for user"
    }
    ```

---

### 3. **Get Image Details**

- **Endpoint**: `GET /images/{imageID}`
- **Method**: `GET`
- **Description**: Retrieves detailed metadata for a single image.
- **Path Parameters**:
  - `imageID` (string) – The ID of the image.
- **Response**:
  - **200 OK**:
    ```json
    {
      "imageId": "string",
      "userId": "string",
      "fileName": "string",
      "uploadDate": "string",
      "lastUpdatedAt": "string",
      "dimensionWidth": "int",
      "dimensionHeight": "int",
      "fileSize": "int64",
      "fileType": "string",
      "analysisStatus": "string",
      "s3Url": "string"
    }
    ```
  - **404 Not Found**: If the image does not exist.
    ```json
    {
      "error": "Image not found"
    }
    ```

---

### 4. **Download Image**

- **Endpoint**: `GET /images/{imageID}/download`
- **Method**: `GET`
- **Description**: Returns the image file for download. (NOTE: Currently not implemented, because we are not uploading file to s3. Its just a json marshalled image meta data. It can be fetched from GET Image Meta API.)
- **Path Parameters**:
  - `imageID` (string) – The ID of the image.
- **Response**:
  - **200 OK**:
    - **Content-Type**: Based on the image type (e.g., `image/jpeg`, `image/png`).
    - **Body**: The raw image file.
  - **404 Not Found**: If the image does not exist.
    ```json
    {
      "error": "Image not found"
    }
    ```

---

### 5. **Update Image Metadata**

- **Endpoint**: `PUT /images/{imageID}`
- **Method**: `PUT`
- **Description**: Updates image metadata such as filename, dimensions, etc.
- **Path Parameters**:
  - `imageID` (string) – The ID of the image.
- **Request Body**:
  ```json
  {
    "fileName": "string",
    "dimensionWidth": "int",
    "dimensionHeight": "int",
    "fileSize": "int64",
    "fileType": "string",
    "analysisStatus": "string",
    "s3Url": "string"
  }
  ```
- **Response**:
  - **200 OK**: Metadata updated successfully.
  - **400 Bad Request**: If the input data is invalid.
  ```json
  {
    "error": "invalid input data"
  }
  ```
  - **404 Not Found**: If the image does not exist.
  ```json
  {
    "error": "Image not found"
  }
  ```

---

### 6. **Delete Image**

- **Endpoint**: `DELETE /images/{imageID}`
- **Method**: `DELETE`
- **Description**: Deletes the image metadata and removes the associated file from S3.
- **Path Parameters**:
  - `imageID` (string) – The ID of the image.
- **Response**:
  - **200 OK**: Image deleted successfully.
  - **404 Not Found**: If the image does not exist.
  ```json
  {
    "error": "Image not found"
  }
  ```
  - **500 Internal Server Error**: If an error occurs while deleting the image.
  ```json
  {
    "error": "Failed to delete image"
  }
  ```

---

## Architecture Diagram

![Arch Diagram](docs/images/Screenshot%202024-11-04%20at%2012.56.52 PM.png)

Currently, I have not implmented Image Analyis Job - which uses reading messages from SQS queue after an image in uploaded or deleted from the S3 bucket of images.

Also, if in future client wants to upload image file. We will send a preSignedURL for upload to the bucket directly by client from upload API. Currently upload API reads data sent in the API.

And once the image analysis job has read the message - it will update the image metadata DB with the additional information.
