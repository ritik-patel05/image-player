package v1_public

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/images/upload", UploadImageMetadata)
	router.GET("/images/user/:userID", GetAllImagesForUser)
	router.GET("/images/:imageID", GetImage)
	router.GET("/images/:imageID/download", DownloadImage)
	router.PUT("/images/:imageID", UpdateImageMetadata)
	router.DELETE("/images/:imageID", DeleteImage)
}
