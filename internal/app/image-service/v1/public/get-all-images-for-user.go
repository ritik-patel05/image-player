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

func GetAllImagesForUser(c *gin.Context) {
	ctx := c.Request.Context()

	req := request.GetAllImagesForUser{
		UserID: c.GetHeader(constants.X_APP_AUTHORIZED_USERID_HEADER),
	}

	if err := req.Validate(); err != nil {
		fmt.Println("GetAllImagesForUser: invalid request", req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := svc.GetAllImagesForUser(ctx, dto.GetAllImagesForUser{
		UserID: req.UserID,
	})
	if err != nil {
		fmt.Println("GetAllImagesForUser: failed to get all images for user", err)
		utils.PrepareErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
