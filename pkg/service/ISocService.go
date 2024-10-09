package service

import "github.com/SocService/pkg/service/models"

type ISocService interface {
	CreateUser(request models.PersonDto) error
	AddFollow(from string, to string) error
	RemoveFollow(from string, to string) error
	GetFollowings(id string, offset int, limit int) ([]models.PersonDto, error)
	GetFollowers(id string, offset int, limit int) ([]models.PersonDto, error)
}
