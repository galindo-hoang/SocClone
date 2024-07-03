package routes

import (
	"github.com/AuthService/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	var router = gin.New()
	setupRoutes(router)
	router.Run(":3000")
}

func setupRoutes(router *gin.Engine) {

	var v1 = router.Group("/v1")
	{
		v1.POST("/sessions", handlers.Login)
		v1.POST("/users", handlers.Register)
		v1.POST("/users/validate", handlers.Validate)
		v1.PUT("/users", handlers.EditUser)
		v1.PUT("/users/avatar", handlers.UploadImage)
	}
}
