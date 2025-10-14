package repository

import (
	m "go-gin-album/internal/model"
)

type AlbumRepository interface {
	FindAll() ([]m.Album, error)
	FindByID(id *uint) (*m.Album, error)
	DeleteByID(id *uint) error
	CreateAlbum(album m.Album) error
}
