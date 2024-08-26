package http

import (
	"errors"
	repo "github.com/AuthService/pkg/repositories/impl"
	"github.com/AuthService/pkg/services/impl"
	"net/http"

	"github.com/AuthService/pkg/internal/rpc"
	model_http "github.com/AuthService/pkg/services/models"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var request model_http.LoginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	svc := impl.NewAuthServices(&repo.AuthRepository{}, &repo.CachingRepository{})

	res, err := svc.Login(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, model_http.Response[model_http.LogInResponse]{
		Success: false,
		Message: "",
		Data:    *res,
	})
}

func Register(ctx *gin.Context) {
	var request model_http.SignUpRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	var caching = repo.CachingRepository{}
	svc := impl.NewAuthServices(&repo.AuthRepository{}, &caching)

	_, isExist := caching.GetCache("register", request.UserName)
	if isExist {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: "user_name already exists",
			Data:    errors.New("user_name already exists"),
		})
		return
	}
	res, err := svc.CreateUser(request)
	if err != nil {
		ctx.JSON(http.StatusOK, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, model_http.Response[model_http.RegisterResponse]{
		Success: false,
		Message: "",
		Data:    *res,
	})
}

func Validate(ctx *gin.Context) {
	var request model_http.ValidateUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	svc := impl.NewAuthServices(&repo.AuthRepository{}, &repo.CachingRepository{})

	res, err := svc.ValidateSigUnUser(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	if err := rpc.CreateNode(*res); err != nil {
		go retriesCreateNode(3, *res)
	}
	ctx.JSON(http.StatusOK, model_http.Response[model_http.RegisterResponse]{
		Success: false,
		Message: "",
		Data:    *res,
	})
}

func retriesCreateNode(times int, node model_http.RegisterResponse) {
	for times > 0 {
		if err := rpc.CreateNode(node); err != nil {
			times--
		} else {
			times = 0
		}
	}
}

func EditUser(ctx *gin.Context) {

}

func UploadImage(ctx *gin.Context) {

}
