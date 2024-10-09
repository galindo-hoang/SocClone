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

	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
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
	id, err := strconv.Atoi(ctx.Param("postId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: true,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
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
	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
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

	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
	err := svc.DeletePost(post.PostId)
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

	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
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

func removeComment(ctx *gin.Context) {
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

	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
	if err := svc.DeleteComment(post); err != nil {
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

	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
	likes, err := svc.LikePost(post)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response[models.LikePostObject]{
		Success: true,
		Data: models.LikePostObject{
			NumLikes: likes,
		},
	})
}

func disLikePost(ctx *gin.Context) {
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

	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
	if err := svc.UnLikePost(post); err != nil {
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
	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
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
	svc := service.NewPostService(repo.NewPostRepository(), repo.NewCloudStorageRepository())
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
