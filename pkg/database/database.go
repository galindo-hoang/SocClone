package database

import (
	"fmt"
	"github.com/PostService/pkg/repositories/models"
	"github.com/PostService/pkg/utils"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/minio/minio-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbName string) error {
	var (
		databaseUser     = utils.GetValue("DB_USER")
		databasePassword = utils.GetValue("DB_PASSWORD")
		databaseHost     = utils.GetValue("DB_HOST")
		databasePort     = utils.GetValue("DB_PORT")
		databaseName     = dbName
	)

	var dataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	database, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := database.AutoMigrate(&models.Posts{}, &models.Images{}, &models.Comments{}, &models.Likes{}); err != nil {
		return err
	}
	DB = database
	fmt.Printf("Connected to database (%v)\n", dbName)
	return nil
}

var Cache *memcache.Client

func InitCache() {
	var (
		cacheHost = utils.GetValue("CACHE_HOST")
		cachePort = utils.GetValue("CACHE_PORT")
	)
	Cache = memcache.New(fmt.Sprintf("%v:%v", cacheHost, cachePort))
	fmt.Printf("Cache: %v\n", Cache)
}

func InitMinio(bucketName string) (*minio.Client, error) {
	var (
		location        = "us-east-1"
		host            = utils.GetValue("MINIO_HOST")
		port            = utils.GetValue("MINIO_PORT")
		accessKeyID     = utils.GetValue("MINIO_ACCESS_KEY")
		secretAccessKey = utils.GetValue("MINIO_SECRET_ACCESS_KEY")
	)
	minioClient, err := minio.New(fmt.Sprintf("%v:%v", host, port), accessKeyID, secretAccessKey, false)
	if err != nil {
		return nil, err
	}

	if err := minioClient.MakeBucket(bucketName, location); err != nil {
		exist, err := minioClient.BucketExists(bucketName)
		if err != nil || !exist {
			return nil, err
		}
	}
	return minioClient, nil
}
