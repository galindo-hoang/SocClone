package http

import (
	model_http "github.com/SocService/pkg/handlers/http/models"
	"github.com/SocService/pkg/internal/grpc"
	repository "github.com/SocService/pkg/repositories/impl"
	"github.com/SocService/pkg/service"
	"github.com/SocService/pkg/service/impl"
	"github.com/SocService/pkg/service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func getListFollow(gin *gin.Context) {
	var socService service.ISocService = impl.NewSocService(&repository.SocRepository{})

	id := gin.DefaultQuery("id", "")
	offset, err := strconv.Atoi(gin.DefaultQuery("offset", "0"))
	if err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	limit, err := strconv.Atoi(gin.DefaultQuery("limit", "20"))

	if err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	res, err := socService.GetFollowings(id, offset, limit)
	if err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	gin.JSON(http.StatusOK, model_http.Response[[]models.PersonDto]{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func follow(gin *gin.Context) {
	var socService service.ISocService = impl.NewSocService(&repository.SocRepository{})

	var req models.FollowDto
	var token = strings.Split(gin.GetHeader("Authorization"), " ")
	if len(token) != 2 {
		gin.JSON(http.StatusUnauthorized, model_http.Response[any]{
			Success: false,
			Message: "Authorization Header is invalid",
			Data:    nil,
		})
		return
	}

	if err := gin.Bind(&req); err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	// check valid token with grpc in authService
	if err := grpc.CheckAuth(req.From, token[1]); err != nil {
		gin.JSON(http.StatusUnauthorized, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
	}
	//
	if err := socService.AddFollow(req.From, req.To); err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	gin.JSON(http.StatusOK, model_http.Response[any]{
		Success: true,
		Message: "",
		Data:    nil,
	})
}

func unfollow(gin *gin.Context) {
	var socService service.ISocService = impl.NewSocService(&repository.SocRepository{})

	var req models.FollowDto
	var token = strings.Split(gin.GetHeader("Authorization"), " ")
	if len(token) != 2 {
		gin.JSON(http.StatusUnauthorized, model_http.Response[any]{
			Success: false,
			Message: "Authorization Header is invalid",
			Data:    nil,
		})
		return
	}

	if err := gin.Bind(&req); err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	// check valid token with grpc in authService

	if err := grpc.CheckAuth(req.From, token[1]); err != nil {
		gin.JSON(http.StatusUnauthorized, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
	}
	//
	if err := socService.RemoveFollow(req.From, req.To); err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	gin.JSON(http.StatusOK, model_http.Response[any]{
		Success: true,
		Message: "",
		Data:    nil,
	})
}

func getListFollower(gin *gin.Context) {
	var socService service.ISocService = impl.NewSocService(&repository.SocRepository{})

	id := gin.DefaultQuery("id", "")
	offset, err := strconv.Atoi(gin.DefaultQuery("offset", "0"))
	if err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	limit, err := strconv.Atoi(gin.DefaultQuery("limit", "20"))

	if err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	res, err := socService.GetFollowers(id, offset, limit)
	if err != nil {
		gin.JSON(http.StatusBadRequest, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	gin.JSON(http.StatusOK, model_http.Response[[]models.PersonDto]{
		Success: true,
		Message: "",
		Data:    res,
	})
}
