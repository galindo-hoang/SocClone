package http

import (
	"errors"
	model_http "github.com/AuthService/pkg/models/http"
	"net/http"

	"github.com/AuthService/pkg/services"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var request model_http.LoginRequest
	ctx.BindJSON(&request)
	res, err := services.Login(request)
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
		Data:    res,
	})
}

func Register(ctx *gin.Context) {
	var request model_http.SignUpRequest
	ctx.BindJSON(&request)
	_, isExist := services.GetCache("register", request.UserName)
	if isExist {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: "user_name already exists",
			Data:    errors.New("user_name already exists"),
		})
		return
	}
	res, err := services.CreateUser(request)
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
		Data:    res,
	})
}

func Validate(ctx *gin.Context) {
	var request model_http.ValidateUserRequest
	ctx.Bind(&request)
	res, err := services.ValidateSigUnUser(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model_http.Response[any]{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, model_http.Response[model_http.RegisterResponse]{
		Success: false,
		Message: "",
		Data:    res,
	})
}

func EditUser(ctx *gin.Context) {

}

func UploadImage(ctx *gin.Context) {

}
