package database

import (
	"fmt"
	"github.com/AuthService/pkg/utils"

	"github.com/AuthService/pkg/repositories/models"
	"github.com/bradfitz/gomemcache/memcache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// check connection to database
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
	if err := database.AutoMigrate(&models.Users{}); err != nil {
		return err
	}
	DB = database
	fmt.Printf("Connected to database (%v)\n", dbName)
	return nil
}

var Cache *memcache.Client

// check connection to server
func InitCache() {
	var (
		cacheHost = utils.GetValue("CACHE_HOST")
		cachePort = utils.GetValue("CACHE_PORT")
	)
	fmt.Printf("hello : %s\n", cacheHost)
	Cache = memcache.New(fmt.Sprintf("%v:%v", cacheHost, cachePort))
	fmt.Printf("Cache: %v\n", Cache)
}
