package service

import (
	"errors"
	message "go-gin-album/internal/constant"
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

func (s *AlbumService) GetAlbumByID(id *uint) (*m.Album, error) {
	albums, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if albums == nil {
		return nil, errors.New(message.AlbumNotFound)
	}

	return albums, nil
}

func (s *AlbumService) DeleteAlbumById(id *uint) (string, error) {
	err := s.Repo.DeleteByID(id)
	if err != nil {
		return "", err
	}

	return message.AlbumDeleted, nil
}

func (s *AlbumService) AddAlbum(newAlbum m.Album) (string, error) {
	if newAlbum.Title == "" {
		return "", errors.New(message.AlbumTitleEmpty)
	}

	err := s.Repo.CreateAlbum(newAlbum)
	if err != nil {
		return "", err
	}

	return message.AlbumAdded, nil
}
