package repositories

import (
	"errors"
	"github.com/AuthService/pkg/database"
	"github.com/AuthService/pkg/repositories/models"
)

type IAuthRepository interface {
	GetUserFrom(username string) (*models.Users, error)
	CreateUser(user models.Users) (*models.Users, error)
	UpdateUser(user models.Users) (*models.Users, error)
}

type AuthRepository struct {
}

func (s *AuthRepository) GetUserFrom(username string) (*models.Users, error) {
	var entity *models.Users = nil
	result := database.DB.Raw("select * from users where user_name = ?", username).Scan(&entity)
	if result.Error != nil {
		return &models.Users{}, result.Error
	}
	if entity == nil {
		return nil, errors.New("user_name invalid")
	}
	return entity, nil
}

func (s *AuthRepository) CreateUser(user models.Users) (*models.Users, error) {
	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	user.Password = ""
	return &user, nil
}

func (s *AuthRepository) UpdateUser(user models.Users) (*models.Users, error) {
	result := database.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
