package http

import (
	"github.com/PostService/pkg/internal/rpc"
	repo "github.com/PostService/pkg/repositories/impl"
	service "github.com/PostService/pkg/services/impl"
	"github.com/PostService/pkg/services/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func createPost(ctx *gin.Context) {
	var post models.PostObject
	token := ctx.GetHeader("Authorization")
	//userAgent := ctx.GetHeader("User-Agent")
	err := ctx.Bind(&post)
	if err != nil {
		log.Println(err.Error())
	}

	if err := rpc.VerifyToken(token, post.Author); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	svc := service.NewPostService(repo.NewPostRepository())
	res, err := svc.CreatePost(post)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response[models.PostRes]{
		Success: true,
		Data:    *res,
	})
}

func getPost(ctx *gin.Context) {
	id := ctx.Param("id")
	svc := service.NewPostService(repo.NewPostRepository())
	res, err := svc.GetPost(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: true,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response[models.PostRes]{
		Success: true,
		Data:    *res,
	})
}

func updatePost(ctx *gin.Context) {
	var post models.MPostObject
	if err := ctx.BindJSON(&post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	token := ctx.GetHeader("Authorization")
	if err := rpc.VerifyToken(token, post.Author); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	svc := service.NewPostService(repo.NewPostRepository())
	res, err := svc.UpdatePost(post)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: true,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response[models.PostRes]{
		Success: true,
		Data:    *res})
}

func deletePost(ctx *gin.Context) {
	var post models.MPostObject
	if err := ctx.BindJSON(&post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	token := ctx.GetHeader("Authorization")
	if err := rpc.VerifyToken(token, post.Author); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	svc := service.NewPostService(repo.NewPostRepository())
	res, err := svc.DeletePost(post.PostId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: true,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response[models.PostRes]{
		Success: true,
		Data:    *res,
	})
}

func commentPost(ctx *gin.Context) {
	var post models.CommentPostObject
	if err := ctx.BindJSON(&post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	token := ctx.GetHeader("Authorization")
	if err := rpc.VerifyToken(token, post.From); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	svc := service.NewPostService(repo.NewPostRepository())
	if err := svc.CommentPost(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response[models.PostRes]{
		Success: true,
	})
}

func likePost(ctx *gin.Context) {
	var post models.LikePostObject
	if err := ctx.BindJSON(&post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	token := ctx.GetHeader("Authorization")
	if err := rpc.VerifyToken(token, post.From); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	svc := service.NewPostService(repo.NewPostRepository())
	if err := svc.LikePost(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response[models.PostRes]{
		Success: true,
	})
}

func getListPosts(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	offSet, err := strconv.Atoi(ctx.Param("offset"))
	if err != nil {
		offSet = 0
	}
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		limit = 20
	}
	var listPost = models.ListPostUserObject{
		ID:     userId,
		Offset: offSet,
		Limit:  limit,
	}
	svc := service.NewPostService(repo.NewPostRepository())
	res, err := svc.GetListPostFromWall(listPost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response[[]*models.PostRes]{
		Success: true,
		Data:    res,
	})
}
