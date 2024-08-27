package services

import (
	"errors"
	"fmt"
	"github.com/AuthService/pkg/database"
	"github.com/AuthService/pkg/internal/mq"
	"github.com/AuthService/pkg/repositories"
	"github.com/AuthService/pkg/repositories/models"
	modelhttp "github.com/AuthService/pkg/services/models"
	"github.com/AuthService/pkg/utils"
	"strconv"
	"time"
)

type IAuthServices interface {
	CreateUser(user modelhttp.SignUpRequest) (*modelhttp.RegisterResponse, error)
	Login(user modelhttp.LoginRequest) (*modelhttp.LogInResponse, error)
	ValidateSigUnUser(user modelhttp.ValidateUserRequest) (*modelhttp.RegisterResponse, error)
}

type AuthServices struct {
	authRepo repositories.IAuthRepository
	caching  repositories.ICachingRepository
}

func NewAuthServices(authRepo repositories.IAuthRepository, caching repositories.ICachingRepository) *AuthServices {
	return &AuthServices{authRepo: authRepo, caching: caching}
}

func (s *AuthServices) CreateUser(user modelhttp.SignUpRequest) (*modelhttp.RegisterResponse, error) {
	exist, err := s.authRepo.GetUserFrom(user.UserName)
	if err != nil && exist != nil {
		return nil, err
	}

	if exist != nil {
		return nil, errors.New("user already registered")
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return nil, err
	}

	if err = s.caching.AddCache("register", user.UserName, otp, 300); err != nil {
		return nil, err
	}

	err = mq.SendMessageMail(mq.MailRequest{
		From:        "username@gmail.com",
		FromName:    "username",
		To:          user.Email,
		Data:        fmt.Sprintf("your otp is: %v", otp),
		ContentType: 0,
	})

	if err != nil {
		return nil, err
	}

	return &modelhttp.RegisterResponse{
		UserName: user.UserName,
		Email:    user.Email,
	}, nil
}

func (s *AuthServices) Login(user modelhttp.LoginRequest) (*modelhttp.LogInResponse, error) {
	entity, err := s.authRepo.GetUserFrom(user.UserName)
	if err != nil {
		return nil, err
	}

	if err := utils.ComparePassword(entity.Password, []byte(user.Password)); err != nil {
		return nil, err
	}
	token, err := utils.BuildingToken(*entity)
	if err != nil {
		return nil, err
	}

	entity.LastActiveAt = time.Now()
	entity.IsActive = true
	database.DB.Save(entity)
	return &modelhttp.LogInResponse{
		ID:       entity.ID,
		Email:    entity.Email,
		Image:    entity.Image,
		UserName: entity.UserName,
		Token:    token,
	}, nil
}

func (s *AuthServices) ValidateSigUnUser(user modelhttp.ValidateUserRequest) (*modelhttp.RegisterResponse, error) {
	item, ok := s.caching.GetCache("register", user.UserName)
	if !ok {
		return nil, errors.New("otp is invalid")
	}
	otp := string(item)
	if otp != user.OTP {
		return nil, errors.New("otp is invalid")
	}

	hashedPassword, err := utils.HashAndSalt([]byte(user.Password))
	if err != nil {
		return nil, err
	}
	entity := &models.Users{
		UserName:     user.UserName,
		Password:     hashedPassword,
		Email:        user.Email,
		CreateAt:     time.Now(),
		LastActiveAt: time.Now(),
	}

	if err := s.caching.DeleteCache("register", user.UserName); err != nil {
		return nil, err
	}

	result := database.DB.Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &modelhttp.RegisterResponse{
		UserName: entity.UserName,
		Email:    entity.Email,
		Id:       strconv.Itoa(entity.ID),
	}, nil
}
