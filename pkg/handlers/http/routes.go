package http

import (
	model_http "github.com/SocService/pkg/handlers/http/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRoutes() {
	var router = gin.New()
	setupRoutes(router)
	log.Println("listening on :3001")
	if err := router.Run(":3001"); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(router *gin.Engine) {
	var v1 = router.Group("/v1")
	{
		v1.GET("/friends/:user_id/*action", func(context *gin.Context) {
			action := context.Param("action")
			if action == "follow" {
				getListFollow(context)
			} else if action == "follower" {
				getListFollower(context)
			} else {
				context.JSON(http.StatusNotFound, model_http.Response[any]{
					Success: false,
					Message: "path not found",
					Data:    nil,
				})
			}
		})
		v1.POST("/friends", follow)
		v1.DELETE("/friends", unfollow)
	}
}
