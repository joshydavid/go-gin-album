package main

import (
	"go-gin/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run()
}
