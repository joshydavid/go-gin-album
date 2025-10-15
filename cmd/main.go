package main

import (
	"context"
	"database/sql"
	"go-gin-album/api"
	"go-gin-album/internal/config"
	"go-gin-album/internal/db"
	"go-gin-album/internal/server"
	"go-gin-album/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// @title           Music Album API
// @version         1.0
// @description     API documentation for Music Album service
// @termsOfService  http://swagger.io/terms/

// @contact.name   Joshua
// @contact.url    https://joshydavid.com
// @contact.email  joshuadavidang@outlook.sg

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	ctx := context.Background()
	rdb := config.SetUpRedisClient(ctx)
	defer rdb.Close()

	albumServiceDependency, sqlDB := initializeDependencies(rdb)
	defer sqlDB.Close()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	api.SetupRoutes(router, albumServiceDependency)
	api.SetUpAPIDocs(router)
	server.RunServer(router)
}

func initializeDependencies(rdb *redis.Client) (*service.AlbumService, *sql.DB) {
	dbConfig := config.LoadDBConfig()
	gormDB, err := db.ConnectDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", err)
	}
	albumService := db.InitializeServices(gormDB, rdb)

	return albumService, sqlDB
}
