package db

import (
	"go-gin-album/internal/repository"
	"go-gin-album/internal/service"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitializeServices(db *gorm.DB, rdb *redis.Client) *service.AlbumService {
	albumRepo := repository.NewPostgresAlbumRepository(db)
	albumService := service.NewAlbumService(albumRepo, rdb)
	return albumService
}
