package dto

import (
	m "go-gin-album/internal/model"
)

type AlbumResponse struct {
	ID     uint    `json:"id" example:"1"`
	Title  string  `json:"title" example:"Blue Train"`
	Artist string  `json:"artist" example:"Joshua David"`
	Price  float64 `json:"price" example:"39.99"`
}

func MapModelToResponse(a m.Album) AlbumResponse {
	return AlbumResponse{
		ID:     a.ID,
		Title:  a.Title,
		Artist: a.Artist,
		Price:  a.Price,
	}
}

func ToResponseSlice(albums []m.Album) []AlbumResponse {
	responses := make([]AlbumResponse, len(albums))
	for i, album := range albums {
		responses[i] = MapModelToResponse(album)
	}
	return responses
}
