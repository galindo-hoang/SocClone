package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func GetValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return os.Getenv(key)
}

func SaveFileToDes(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(src)

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	_, err = io.Copy(out, src)
	return err
}

func DeleteDirectory(path string) {
	if err := os.RemoveAll(path); err != nil {
		log.Fatal(err)
	}
}

func GetPathFromMinio(bucket string, image string) string {
	var (
		host = GetValue("MINIO_HOST")
		port = GetValue("MINIO_PORT_WEB")
		path = "api/v1"
	)
	var res = fmt.Sprintf("http://%s:%s/%s", host, port, path)
	if len(bucket) != 0 {
		res = fmt.Sprintf("%s/buckets/%s", res, bucket)
	}
	res = fmt.Sprintf("%s/objects/download?preview=true&prefix=%s", res, image)
	return res
}
