package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/AuthService/pkg/interal/mq"
	modelhttp "github.com/AuthService/pkg/models/http"
	"github.com/AuthService/pkg/models/rbmq"

	"github.com/AuthService/pkg/database"
	"github.com/AuthService/pkg/models"
	"github.com/AuthService/utils"
	"github.com/bradfitz/gomemcache/memcache"
)

func CreateUser(user modelhttp.SignUpRequest) (modelhttp.RegisterResponse, error) {
	var entity *models.Users = nil
	database.DB.Raw("select * from users where user_name = ?", user.UserName).Scan(&entity)
	if entity != nil {
		return modelhttp.RegisterResponse{}, errors.New("user already registered")
	}
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return modelhttp.RegisterResponse{}, err
	}

	err = mq.SendMessageMail(rbmq.MailRequest{
		From:        "username@gmail.com",
		FromName:    "username",
		To:          user.Email,
		Data:        fmt.Sprintf("your otp is: %v", otp),
		ContentType: 0,
	})

	if err != nil {
		return modelhttp.RegisterResponse{}, err
	}
	addNewCache("register", user.UserName, otp, 300)

	return modelhttp.RegisterResponse{
		UserName: user.UserName,
		Email:    user.Email,
	}, nil
}

func Login(user modelhttp.LoginRequest) (modelhttp.LogInResponse, error) {
	var entity *models.Users = nil
	database.DB.Raw("select * from users where user_name = ?", user.UserName).Scan(&entity)

	if entity == nil {
		return modelhttp.LogInResponse{}, errors.New("user_name or password invalid")
	}

	if err := utils.ComparePassword(entity.Password, []byte(user.Password)); err != nil {
		return modelhttp.LogInResponse{}, err
	}
	token, err := utils.BuildingToken(*entity)
	if err != nil {
		return modelhttp.LogInResponse{}, err
	}

	entity.LastActiveAt = time.Now()
	entity.IsActive = true
	database.DB.Save(entity)
	return modelhttp.LogInResponse{
		ID:       entity.ID,
		Email:    entity.Email,
		Image:    entity.Image,
		UserName: entity.UserName,
		Token:    token,
	}, nil
}

func ValidateSigUnUser(user modelhttp.ValidateUserRequest) (modelhttp.RegisterResponse, error) {
	item, ok := GetCache("register", user.UserName)
	if !ok {
		return modelhttp.RegisterResponse{}, errors.New("otp is invalid")
	}
	otp := string(item)
	if otp != user.OTP {
		return modelhttp.RegisterResponse{}, errors.New("otp is invalid")
	}

	hashedPassword, err := utils.HashAndSalt([]byte(user.Password))
	if err != nil {
		return modelhttp.RegisterResponse{}, err
	}
	entity := &models.Users{
		UserName:     user.UserName,
		Password:     hashedPassword,
		Email:        user.Email,
		CreateAt:     time.Now(),
		LastActiveAt: time.Now(),
	}

	if err := DeleteCache("register", user.UserName); err != nil {
		return modelhttp.RegisterResponse{}, err
	}

	result := database.DB.Create(entity)
	if result.Error != nil {
		return modelhttp.RegisterResponse{}, result.Error
	}

	return modelhttp.RegisterResponse{
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
