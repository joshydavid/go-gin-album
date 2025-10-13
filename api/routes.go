package api

import (
	"go-gin-album/internal/constant"
	"go-gin-album/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group(constant.BasePath)
	{
		v1.GET(constant.HealthCheck, handler.GetHealthCheck)
		v1.GET(constant.Albums, handler.GetAlbums)
		v1.GET(constant.AlbumByID, handler.GetAlbumByID)
		v1.POST(constant.Albums, handler.AddAlbum)
		v1.DELETE(constant.AlbumByID, handler.DeleteAlbumByID)
	}
}
