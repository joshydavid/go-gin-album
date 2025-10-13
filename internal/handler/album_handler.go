package handler

import (
	"net/http"
	"strconv"
	"strings"

	"go-gin-album/internal/dto"
	"go-gin-album/internal/model"
	"go-gin-album/internal/service"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	AlbumService *service.AlbumService
}

func NewAlbumHandler(s *service.AlbumService) *AlbumHandler {
	return &AlbumHandler{AlbumService: s}
}

// GetAllAlbums godoc (Renamed from GetAlbums for clarity/convention)
// @Summary List all albums
// @Schemes
// @Description Retrieves a list of all music albums from the collection.
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} dto.AlbumResponse
// @Failure	500 {object} map[string]string "Internal Server Error"
// @Router /albums [get]
func (h *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := h.AlbumService.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := dto.ToResponseSlice(albums)
	c.IndentedJSON(http.StatusOK, response)
}

// GetAlbumByID godoc
// @Summary Get album by ID
// @Schemes
// @Description Retrieve a single music album by its unique ID.
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} dto.AlbumResponse
// @Failure 404 {object} map[string]string "Album not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /albums/{id} [get]
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := h.AlbumService.GetAlbumByID(id)

	if err != nil {
		if strings.Contains(err.Error(), "album not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := dto.MapModelToResponse(*album)
	c.IndentedJSON(http.StatusOK, response)
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
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /albums/{id} [delete]
func (h *AlbumHandler) DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID format. Must be a positive integer."})
		return
	}

	uintID := uint(parsedID)
	_, err = h.AlbumService.DeleteAlbumById(&uintID)

	if err != nil {
		if strings.Contains(err.Error(), "album not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Param album body dto.AlbumResponse true "Album object to be created"
// @Success 201 {object} dto.AlbumResponse
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /albums [post]
func (h *AlbumHandler) AddAlbum(c *gin.Context) {
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.AlbumService.AddAlbum(newAlbum)

	if err != nil {
		if strings.Contains(err.Error(), "cannot be empty") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.MapModelToResponse(newAlbum)
	c.JSON(http.StatusCreated, response)
}
