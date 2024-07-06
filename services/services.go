package services

import (
	"errors"
	"fmt"

	"github.com/AuthService/database"
	"github.com/AuthService/models"
	"github.com/AuthService/utils"
	"github.com/bradfitz/gomemcache/memcache"
)

func CraeteUser(user models.SignUpRequest) (models.RegisterResponse, error) {
	var entity *models.Users = nil
	database.DB.Raw("select * from users where user_name = ?", user.UserName).Scan(&entity)
	if entity != nil {
		return models.RegisterResponse{}, errors.New("user already registered")
	}
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return models.RegisterResponse{}, err
	}
	addNewCache("register", user.UserName, otp, 300)
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

func ValidateSigUnUser(user models.ValidateUserRequest) (models.RegisterResponse, error) {
	item, ok := GetCache("register", user.UserName)
	if !ok {
		return models.RegisterResponse{}, errors.New("otp is invalid")
	}
	otp, err := utils.ToJsonFromByte[string](item)
	if err != nil || otp != user.OTP {
		return models.RegisterResponse{}, errors.New("otp is invalid")
	}

	hashedPassword, err := utils.HashAndSalt([]byte(user.Password))
	if err != nil {
		return models.RegisterResponse{}, err
	}
	entity := &models.Users{
		UserName: user.UserName,
		Password: hashedPassword,
		Email:    user.Email,
	}

	if err := DeleteCache("register", user.UserName); err != nil {
		return models.RegisterResponse{}, err
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

func GetCache(path string, key string) ([]byte, bool) {
	item, err := database.Cache.Get(fmt.Sprintf("%v/%v", path, key))
	if err == nil {
		return item.Value, true
	}
	return []byte{}, false
}

func addNewCache(path string, key string, value string, time int) error {
	if time != 0 {
		cacheError := database.Cache.Set(&memcache.Item{
			Key:        fmt.Sprintf("%v/%v", path, key),
			Value:      []byte(value),
			Expiration: int32(time),
		})
		if cacheError != nil {
			return cacheError
		}
	} else {
		cacheError := database.Cache.Set(&memcache.Item{
			Key:   fmt.Sprintf("%v/%v", path, key),
			Value: []byte(value),
		})
		if cacheError != nil {
			return cacheError
		}
	}
	return nil
}

func DeleteCache(path string, key string) error {
	err := database.Cache.Delete(fmt.Sprintf("%v/%v", path, key))
	return err
}
