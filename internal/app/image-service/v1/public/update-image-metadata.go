package v1_public

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritik-patel05/image-player/internal/constants"
	svc "github.com/ritik-patel05/image-player/internal/domain/service/image-service"
	"github.com/ritik-patel05/image-player/internal/models/dto"
	"github.com/ritik-patel05/image-player/internal/models/request"
	"github.com/ritik-patel05/image-player/internal/utils"
)

func UpdateImageMetadata(c *gin.Context) {
	ctx := c.Request.Context()

	req := request.UpdateImageMetadata{
		UserID:  c.GetHeader(constants.X_APP_AUTHORIZED_USERID_HEADER),
		ImageID: c.Param("imageID"),
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("UpdateImageMetadata: error while parsing req body", req)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := req.Validate(); err != nil {
		fmt.Println("UpdateImageMetadata: invalid request", req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := svc.UpdateImageMetadata(ctx, dto.UpdateImageMetadata{
		ImageID:         req.ImageID,
		FileName:        req.FileName,
		FileSize:        req.FileSize,
		FileType:        req.FileType,
		DimensionWidth:  req.DimensionWidth,
		DimensionHeight: req.DimensionHeight,
		AnalysisStatus:  req.AnalysisStatus,
		S3URL:           req.S3URL,
	})
	if err != nil {
		fmt.Println("UpdateImageMetadata: failed to get image", err)
		utils.PrepareErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
