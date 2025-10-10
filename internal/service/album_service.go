package service

import (
	"errors"
	"go-gin/internal/model"
)

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllAlbums() ([]model.Album, error) {
	return albums, nil
}

func GetAlbumByID(id string) (model.Album, error) {
	for _, a := range albums {
		if a.ID == id {
			return a, nil
		}
	}
	return model.Album{}, errors.New("album not found")
}

func AddAlbum(newAlbum model.Album) (string, error) {
	if newAlbum.Title == "" {
		return "", errors.New("album title cannot be empty")
	}

	albums = append(albums, newAlbum)
	return "Album added successfully", nil
}
