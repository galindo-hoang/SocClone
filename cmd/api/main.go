package main

import (
	"github.com/AuthService/pkg/database"
	"github.com/AuthService/pkg/handlers/http"
	"github.com/AuthService/pkg/handlers/rpc"
	"github.com/AuthService/pkg/utils"
)

func main() {
	if err := database.InitDatabase(utils.GetValue("DB_NAME")); err != nil {
		panic(err)
	}
	database.InitCache()
	go rpc.NewAuthHandler()
	http.InitRoutes()
}
