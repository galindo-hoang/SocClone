package repositories

import "github.com/AuthService/pkg/repositories/models"

type IAuthRepository interface {
	GetUserFrom(username string) (*models.Users, error)
	CreateUser(user models.Users) (*models.Users, error)
	UpdateUser(user models.Users) (*models.Users, error)
	//DeleteUser(username string, pwd string) error
}
