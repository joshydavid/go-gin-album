package api

import (
	"go-gin/internal/constant"
	"go-gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET(constant.HealthCheck, handler.GetHealthCheck)
	router.GET(constant.Albums, handler.GetAlbums)
}
