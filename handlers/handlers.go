package handlers

import (
	"errors"
	"net/http"

	"github.com/AuthService/models"
	"github.com/AuthService/services"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var request models.LoginRequest
	ctx.BindJSON(&request)
	res, err := services.Login(request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusAccepted, res)
}

func Register(ctx *gin.Context) {
	var request models.SignUpRequest
	ctx.BindJSON(&request)
	_, isExist := services.GetCache("register", request.UserName)
	if isExist {
		ctx.JSON(http.StatusOK, errors.New("user_name already exists"))
		return
	}
	res, err := services.CraeteUser(request)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func Validate(ctx *gin.Context) {
	var request models.ValidateUserRequest
	ctx.Bind(&request)
	res, err := services.ValidateSigUnUser(request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func EditUser(ctx *gin.Context) {

}

func UploadImage(ctx *gin.Context) {

}
