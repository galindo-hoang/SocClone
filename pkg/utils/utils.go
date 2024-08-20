package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return os.Getenv(key)
}
