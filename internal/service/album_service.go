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
	Repo        repository.AlbumRepository
	RedisClient *redis.Client
}

func NewAlbumService(repo repository.AlbumRepository, rdb *redis.Client) *AlbumService {
	return &AlbumService{
		Repo:        repo,
		RedisClient: rdb,
	}
}

func (s *AlbumService) GetAllAlbums() ([]m.Album, error) {
	ctx := context.Background()
	cachedData, err := s.RedisClient.Get(ctx, c.AllAlbumsCacheKey).Bytes()
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
		s.RedisClient.Set(ctx, c.AllAlbumsCacheKey, jsonAlbums, c.DefaultTTLMinutes*time.Minute)
	}

	return albums, nil
}

func (s *AlbumService) GetAlbumByID(id *uint) (*m.Album, error) {
	if id == nil {
		return nil, errors.New(c.InvalidAlbumID)
	}

	ctx := context.Background()
	albumKey := fmt.Sprintf(c.AlbumCacheKey, *id)
	cachedData, err := s.RedisClient.Get(ctx, albumKey).Bytes()
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
		s.RedisClient.Set(ctx, albumKey, jsonAlbum, c.DefaultTTLMinutes*time.Minute)
	}

	return album, nil
}

func (s *AlbumService) DeleteAlbumById(id *uint) (string, error) {
	err := s.Repo.DeleteByID(id)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	albumKey := fmt.Sprintf(c.AlbumCacheKey, *id)
	s.RedisClient.Del(ctx, albumKey)
	s.RedisClient.Del(ctx, c.AllAlbumsCacheKey)

	return c.AlbumDeleted, nil
}

func (s *AlbumService) AddAlbum(newAlbum m.Album) (string, error) {
	if newAlbum.Title == "" {
		return "", errors.New(c.AlbumTitleEmpty)
	}

	err := s.Repo.CreateAlbum(newAlbum)
	if err != nil {
		return "", err
	}

	s.RedisClient.Del(context.Background(), c.AllAlbumsCacheKey)

	return c.AlbumAdded, nil
}
