package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v1_public "github.com/ritik-patel05/image-player/internal/app/image-service/v1/public"
	"github.com/ritik-patel05/image-player/internal/utils"
)

func NewServer() *gin.Engine {
	if utils.IsProductionEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	}))

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Image Service")
	})

	v1RouterPublic := router.Group("/image-service/public/v1")
	v1_public.RegisterRoutes(v1RouterPublic)

	return router
}

func Shutdown(ctx context.Context, srv *http.Server) {
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Error in Shutdown", err)
	}
}
