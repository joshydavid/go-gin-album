package api

import (
	"go-gin-album/internal/constant"
	"go-gin-album/internal/handler"
	"go-gin-album/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, albumService *service.AlbumService, rateLimitMiddleware gin.HandlerFunc) {
	albumHandler := handler.NewAlbumHandler(albumService)
	v1 := router.Group(constant.BasePath)
	v1.Use(rateLimitMiddleware)
	{
		v1.GET(constant.HealthCheck, handler.GetHealthCheck)
		v1.GET(constant.Albums, albumHandler.GetAllAlbums)
		v1.GET(constant.AlbumByID, albumHandler.GetAlbumByID)
		v1.POST(constant.Albums, albumHandler.AddAlbum)
		v1.DELETE(constant.AlbumByID, albumHandler.DeleteAlbumByID)
	}
}
