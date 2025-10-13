package handler

import (
	"go-gin-album/internal/model"
	"go-gin-album/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAlbums godoc
// @Summary List all albums
// @Schemes
// @Description Retrieves a list of all music albums from the collection.
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} model.Album
// @Failure	500 {object} map[string]string "Internal Server Error"
// @Router /albums [get]
func GetAlbums(c *gin.Context) {
	albums, err := service.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID godoc
// @Summary Get album by ID
// @Schemes
// @Description Retrieve a single music album by its unique ID.
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} model.Album
// @Failure 404 {object} map[string]string "Album not found"
// @Router /albums/{id} [get]
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := service.GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

// DeleteAlbumByID godoc
// @Summary Delete an album
// @Schemes
// @Description Deletes a single music album record by its unique ID.
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 204 {object} map[string]string "No Content"
// @Failure 404 {object} map[string]string "Album not found"
// @Router /albums/{id} [delete]
func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	_, err := service.DeleteAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// AddAlbum godoc
// @Summary Add a new album
// @Schemes
// @Description Creates and stores a new music album record.
// @Tags albums
// @Accept json
// @Produce json
// @Param album body model.Album true "Album object to be created"
// @Success 201 {object} model.Album
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /albums [post]
func AddAlbum(c *gin.Context) {
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAlbum, err := service.AddAlbum(newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdAlbum)
}
