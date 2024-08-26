package impl

import (
	"fmt"
	"github.com/AuthService/pkg/database"
)

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
