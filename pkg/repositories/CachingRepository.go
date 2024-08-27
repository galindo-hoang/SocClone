package repositories

import (
	"fmt"
	"github.com/AuthService/pkg/database"
	"github.com/bradfitz/gomemcache/memcache"
)

type ICachingRepository interface {
	DeleteCache(path string, key string) error
	GetCache(path string, key string) ([]byte, bool)
	AddCache(path string, key string, value string, time int) error
}

type CachingRepository struct {
}

func (s *CachingRepository) DeleteCache(path string, key string) error {
	err := database.Cache.Delete(fmt.Sprintf("%v/%v", path, key))
	return err
}

func (s *CachingRepository) GetCache(path string, key string) ([]byte, bool) {
	item, err := database.Cache.Get(fmt.Sprintf("%v/%v", path, key))
	if err == nil {
		return item.Value, true
	}
	return []byte{}, false
}

func (s *CachingRepository) AddCache(path string, key string, value string, time int) error {
	if time != 0 {
		cacheError := database.Cache.Set(&memcache.Item{
			Key:        fmt.Sprintf("%v/%v", path, key),
			Value:      []byte(value),
			Expiration: int32(time),
		})
		if cacheError != nil {
			return cacheError
		}
	} else {
		cacheError := database.Cache.Set(&memcache.Item{
			Key:   fmt.Sprintf("%v/%v", path, key),
			Value: []byte(value),
		})
		if cacheError != nil {
			return cacheError
		}
	}
	return nil
}
