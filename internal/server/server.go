package server

import (
	"go-gin-album/pkg/util"
	"os"

	"github.com/gin-gonic/gin"
)

func RunServer(router *gin.Engine) {
	util.LoadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
