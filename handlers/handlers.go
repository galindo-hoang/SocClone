package handlers

import (
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

}

func Validate(ctx *gin.Context) {

}

func EditUser(ctx *gin.Context) {

}

func UploadImage(ctx *gin.Context) {

}
