package v1_public

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	svc "github.com/ritik-patel05/image-player/internal/domain/service/image-service"
	"github.com/ritik-patel05/image-player/internal/models/dto"
	"github.com/ritik-patel05/image-player/internal/models/request"
	"github.com/ritik-patel05/image-player/internal/utils"
)

func DownloadImage(c *gin.Context) {
	ctx := c.Request.Context()

	req := request.DownloadImage{
		ImageID: c.Param("imageID"),
	}
	if err := req.Validate(); err != nil {
		fmt.Println("DownloadImage: invalid request", req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := svc.DownloadImage(ctx, dto.DownloadImage{
		ImageID: req.ImageID,
	})
	if err != nil {
		fmt.Println("DownloadImage: failed to download image", err)
		utils.PrepareErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
