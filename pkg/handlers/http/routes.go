package http

import (
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
		v1.POST("/sessions", Login)
		v1.POST("/users", Register)
		v1.POST("/users/validate", Validate)
		v1.PUT("/users", EditUser)
		v1.PUT("/users/avatar", UploadImage)
	}
}
