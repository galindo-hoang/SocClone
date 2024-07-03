package database

import (
	"fmt"

	"github.com/AuthService/models"
	"github.com/AuthService/utils"
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

	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	database, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := database.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	DB = database
	fmt.Printf("Connected to database (%v)\n", dbName)
	return nil
}
