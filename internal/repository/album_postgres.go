package repository

import (
	"errors"
	m "go-gin-album/internal/model"

	"gorm.io/gorm"
)

type PostgresAlbumRepository struct {
	DB *gorm.DB
}

func NewPostgresAlbumRepository(db *gorm.DB) AlbumRepository {
	return &PostgresAlbumRepository{
		DB: db,
	}
}

func (r *PostgresAlbumRepository) FindAll() ([]m.Album, error) {
	var albums []m.Album
	result := r.DB.Find(&albums)
	return albums, result.Error
}

func (r *PostgresAlbumRepository) FindByID(id *uint) (*m.Album, error) {
	var album m.Album
	result := r.DB.First(&album, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &album, nil
}

func (r *PostgresAlbumRepository) DeleteByID(id *uint) error {
	result := r.DB.Delete(&m.Album{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("album not found")
	}

	return nil
}

func (r *PostgresAlbumRepository) CreateAlbum(album m.Album) error {
	result := r.DB.Create(&album)
	return result.Error
}
