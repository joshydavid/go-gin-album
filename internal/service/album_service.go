package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	c "go-gin-album/internal/constant"
	m "go-gin-album/internal/model"
	"go-gin-album/internal/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type AlbumService struct {
	Repo  repository.AlbumRepository
	Cache redis.Cmdable
}

func NewAlbumService(repo repository.AlbumRepository, cache redis.Cmdable) *AlbumService {
	return &AlbumService{
		Repo:  repo,
		Cache: cache,
	}
}

func (s *AlbumService) GetAllAlbums(ctx context.Context) ([]m.Album, error) {
	cachedData, err := s.Cache.Get(ctx, c.AllAlbumsCacheKey).Bytes()
	if err == nil {
		var albums []m.Album
		if err := json.Unmarshal(cachedData, &albums); err == nil {
			return albums, nil
		}
	}

	albums, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	jsonAlbums, err := json.Marshal(albums)
	if err == nil {
		s.Cache.Set(ctx, c.AllAlbumsCacheKey, jsonAlbums, c.DefaultTTLMinutes*time.Minute)
	}

	return albums, nil
}

func (s *AlbumService) GetAlbumByID(ctx context.Context, id *uint) (*m.Album, error) {
	if id == nil {
		return nil, errors.New(c.InvalidAlbumID)
	}

	albumKey := fmt.Sprintf(c.AlbumCacheKey, *id)
	cachedData, err := s.Cache.Get(ctx, albumKey).Bytes()
	if err == nil {
		var album m.Album
		if err := json.Unmarshal(cachedData, &album); err == nil {
			return &album, nil
		}
	}

	album, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if album == nil {
		return nil, errors.New(c.AlbumNotFound)
	}

	jsonAlbum, err := json.Marshal(album)
	if err == nil {
		s.Cache.Set(ctx, albumKey, jsonAlbum, c.DefaultTTLMinutes*time.Minute)
	}

	return album, nil
}

func (s *AlbumService) DeleteAlbumById(ctx context.Context, id *uint) (string, error) {
	err := s.Repo.DeleteByID(id)
	if err != nil {
		return "", err
	}

	albumKey := fmt.Sprintf(c.AlbumCacheKey, *id)
	s.Cache.Del(ctx, albumKey)
	s.Cache.Del(ctx, c.AllAlbumsCacheKey)

	return c.AlbumDeleted, nil
}

func (s *AlbumService) AddAlbum(ctx context.Context, newAlbum m.Album) (string, error) {
	if newAlbum.Title == "" {
		return "", errors.New(c.AlbumTitleEmpty)
	}

	err := s.Repo.CreateAlbum(newAlbum)
	if err != nil {
		return "", err
	}

	s.Cache.Del(ctx, c.AllAlbumsCacheKey)

	return c.AlbumAdded, nil
}
