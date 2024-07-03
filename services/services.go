package services

import (
	"errors"

	"github.com/AuthService/database"
	"github.com/AuthService/models"
	"github.com/AuthService/utils"
)

func CraeteUser(user models.SignUpRequest) (models.RegisterResponse, error) {
	var entity *models.Users = nil
	database.DB.Raw("select * from users where user_name = ?", user.UserName).Scan(&entity)
	if entity != nil {
		return models.RegisterResponse{}, errors.New("user already registered")
	}
	return models.RegisterResponse{
		UserName: user.UserName,
		Email:    user.Email,
	}, nil
}

func Login(user models.LoginRequest) (models.LogInResponse, error) {
	var entity *models.Users = nil
	database.DB.Raw("select * from users where user_name = ?", user.UserName).Scan(&entity)

	if entity == nil {
		return models.LogInResponse{}, errors.New("user_name or password invalid")
	}

	if err := utils.ComparePassword(entity.Password, []byte(user.Password)); err != nil {
		return models.LogInResponse{}, err
	}

	return models.LogInResponse{
		ID:       entity.ID,
		Email:    entity.Email,
		Image:    entity.Image,
		UserName: entity.UserName,
	}, nil
}

func ValidateUser(user models.ValidateUserRequest) (models.RegisterResponse, error) {
	var entity *models.Users = nil
	database.DB.Raw("select * from users where user_name = ?", user.UserName).Scan(&entity)
	if entity != nil {
		return models.RegisterResponse{}, errors.New("user already registered")
	}
	hashedPassword, err := utils.HashAndSalt([]byte(user.Password))
	if err != nil {
		return models.RegisterResponse{}, err
	}
	entity = &models.Users{
		UserName: user.UserName,
		Password: hashedPassword,
		Email:    entity.Email,
	}

	result := database.DB.Create(entity)
	if result.Error != nil {
		return models.RegisterResponse{}, result.Error
	}

	return models.RegisterResponse{
		UserName: entity.UserName,
		Email:    entity.Email,
	}, nil
}
