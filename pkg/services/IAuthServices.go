package services

import modelhttp "github.com/AuthService/pkg/services/models"

type IAuthServices interface {
	CreateUser(user modelhttp.SignUpRequest) (*modelhttp.RegisterResponse, error)
	Login(user modelhttp.LoginRequest) (*modelhttp.LogInResponse, error)
	ValidateSigUnUser(user modelhttp.ValidateUserRequest) (*modelhttp.RegisterResponse, error)
}
