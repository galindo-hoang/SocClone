package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

func New() {
	router := gin.New()
	setupRoutes(router)
	err := router.Run(":3002")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(router *gin.Engine) {
	var v1 = router.Group("/v1")
	{
		v1.GET("/friends/:user_id/posts", getListPosts)
	}

	var posts = v1.Group("/posts")
	{
		posts.POST("/create", createPost)
		posts.GET("/:postId", getPost)
		posts.PUT("/", updatePost)
		posts.DELETE("/", deletePost)

		posts.POST("/comment", commentPost)
		posts.DELETE("/comment", removeComment)
		posts.POST("/likes", likePost)
		posts.DELETE("/likes", disLikePost)
	}
}
