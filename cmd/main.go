package main

import (
	"context"
	"database/sql"
	"go-gin-album/api"
	"go-gin-album/api/middleware"
	"go-gin-album/internal/config"
	"go-gin-album/internal/db"
	"go-gin-album/internal/server"
	"go-gin-album/internal/service"
	"go-gin-album/pkg/util"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
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
	util.LoadEnv()
	ctx := context.Background()
	rdb := config.SetUpRedisClient(ctx)
	defer rdb.Close()

	albumServiceDependency, sqlDB := initializeDb(rdb)
	defer sqlDB.Close()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	rateLimitMiddleware := initializeRateLimiter()
	api.SetupRoutes(router, albumServiceDependency, rateLimitMiddleware)
	api.SetUpAPIDocs(router)
	server.RunServer(router)
}

func initializeDb(rdb *redis.Client) (*service.AlbumService, *sql.DB) {
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

func initializeRateLimiter() gin.HandlerFunc {
	const defaultRate, defaultBurst = 1, 5

	tokenRefillRate, err := strconv.Atoi(os.Getenv("TOKEN_REFILL_RATE"))
	if err != nil {
		log.Printf("Warning: Invalid or unset TOKEN_REFILL_RATE '%d'. Using default rate: %d\n", tokenRefillRate, defaultRate)
		tokenRefillRate = defaultRate
	}

	burstSize, err := strconv.Atoi(os.Getenv("BURST_SIZE"))
	if err != nil {
		log.Printf("Warning: Invalid or unset BURST_SIZE '%d'. Using default burst: %d\n", burstSize, defaultBurst)
		burstSize = defaultBurst
	}

	if tokenRefillRate <= 0 {
		log.Println("Warning: Calculated TOKEN_REFILL_RATE is non-positive. Setting to 1.")
		tokenRefillRate = 1
	}

	rateLimiter := rate.NewLimiter(rate.Limit(tokenRefillRate), burstSize)
	rateLimitMiddleware := middleware.RateLimiter(rateLimiter)

	return rateLimitMiddleware
}
