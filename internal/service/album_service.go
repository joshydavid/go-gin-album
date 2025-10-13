package service

import (
	"errors"
	m "go-gin-album/internal/model"
	"go-gin-album/internal/repository"
)

type AlbumService struct {
	Repo repository.AlbumRepository
}

func NewAlbumService(repo repository.AlbumRepository) *AlbumService {
	return &AlbumService{
		Repo: repo,
	}
}

func (s *AlbumService) GetAllAlbums() ([]m.Album, error) {
	albums, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (s *AlbumService) GetAlbumByID(id string) (*m.Album, error) {
	albums, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("database error when retrieving album: " + err.Error())
	}

	if albums == nil {
		return nil, errors.New("album not found")
	}

	return albums, nil
}

func (s *AlbumService) DeleteAlbumById(id *uint) (string, error) {
	err := s.Repo.DeleteByID(id)
	if err != nil {
		return "", errors.New("database error when deleting album: " + err.Error())
	}

	return "Album deleted", nil
}

func (s *AlbumService) AddAlbum(newAlbum m.Album) (string, error) {
	if newAlbum.Title == "" {
		return "", errors.New("album title cannot be empty")
	}

	err := s.Repo.CreateAlbum(newAlbum)
	if err != nil {
		return "", errors.New("database error when adding album: " + err.Error())
	}

	return "Album added", nil
}
