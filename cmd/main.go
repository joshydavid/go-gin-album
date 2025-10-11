package main

import (
	"go-gin-album/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	api.SetupRoutes(router)
	RunServer(router)
}

func RunServer(router *gin.Engine) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
