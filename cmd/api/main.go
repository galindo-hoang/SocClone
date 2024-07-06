package main

import (
	"github.com/AuthService/database"
	"github.com/AuthService/routes"
	"github.com/AuthService/utils"
)

func main() {
	routes.InitRoutes()

	database.InitCache()
	database.InitDatabase(utils.GetValue("DB_NAME"))
}
