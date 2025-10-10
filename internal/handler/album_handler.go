package handler

import (
	"go-gin/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albums, err := service.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}
