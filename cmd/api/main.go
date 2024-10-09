package main

import (
	"github.com/PostService/pkg/database"
	httproot "github.com/PostService/pkg/handlers/http"
	"github.com/PostService/pkg/utils"
	"log"
)

func main() {
	if err := database.InitDatabase(utils.GetValue("DB_NAME")); err != nil {
		log.Fatal(err)
	}
	httproot.New()
}
