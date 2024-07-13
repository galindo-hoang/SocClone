package main

import (
	"github.com/AuthService/pkg/database"
	"github.com/AuthService/pkg/handlers/http"
	"github.com/AuthService/utils"
)

func main() {
	if err := database.InitDatabase(utils.GetValue("DB_NAME")); err != nil {
		panic(err)
	}
	database.InitCache()
	http.InitRoutes()

}
