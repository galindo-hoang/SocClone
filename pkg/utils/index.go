package utils

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

func GetValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}

	return os.Getenv(key)
}

func Byte2Json[T any](val []byte) (T, error) {
	var data = new(T)
	err := json.Unmarshal(val, &data)
	if err != nil {
		return *data, err
	}
	return *data, nil
}
