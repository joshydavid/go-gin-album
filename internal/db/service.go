package db

import (
	"go-gin-album/internal/repository"
	"go-gin-album/internal/service"

	"gorm.io/gorm"
)

func InitializeServices(db *gorm.DB) *service.AlbumService {
	albumRepo := repository.NewPostgresAlbumRepository(db)
	albumService := service.NewAlbumService(albumRepo)
	return albumService
}
