package main

import (
	"go-gin-album/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run()
}
