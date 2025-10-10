package api

import (
	"go-gin-album/internal/constant"
	"go-gin-album/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET(constant.HealthCheck, handler.GetHealthCheck)
	router.GET(constant.Albums, handler.GetAlbums)
	router.GET(constant.AlbumByID, handler.GetAlbumByID)
	router.POST(constant.Albums, handler.AddAlbum)
	router.DELETE(constant.AlbumByID, handler.DeleteAlbumByID)
}
