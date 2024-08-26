package repositories

type ICachingRepository interface {
	DeleteCache(path string, key string) error
	GetCache(path string, key string) ([]byte, bool)
}
