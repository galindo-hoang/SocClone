package impl

import (
	"errors"
	"github.com/SocService/pkg/repositories"
	model "github.com/SocService/pkg/repositories/model"
	"github.com/SocService/pkg/service/models"
)

type SocService struct {
	socRepo repositories.ISocRepository
}

func NewSocService(socRepo repositories.ISocRepository) *SocService {
	return &SocService{socRepo: socRepo}
}

func (s *SocService) CreateUser(request models.PersonDto) error {
	if len(request.Id) == 0 || len(request.Name) == 0 {
		return errors.New("request is invalid")
	}
	return s.socRepo.CreateNode(model.Person{
		Id:    request.Id,
		Name:  request.Name,
		Image: request.Image,
	})
}
func (s *SocService) AddFollow(from string, to string) error {
	if len(from) == 0 || len(to) == 0 {
		return errors.New("request is invalid")
	}
	return s.socRepo.MakeRelation(from, to)
}

func (s *SocService) RemoveFollow(from string, to string) error {
	if len(from) == 0 || len(to) == 0 {
		return errors.New("request is invalid")
	}
	return s.socRepo.MakeDetach(from, to)
}
func (s *SocService) GetFollowings(id string, offset int, limit int) ([]models.PersonDto, error) {
	if len(id) == 0 {
		return nil, errors.New("request is invalid")
	}
	persons, err := s.socRepo.GetFollowings(id, offset, limit)
	if err != nil {
		return nil, err
	}
	var result []models.PersonDto
	for _, person := range persons {
		result = append(result, models.PersonDto{
			Id:    person.Id,
			Name:  person.Name,
			Image: person.Image,
		})
	}
	return result, nil
}
func (s *SocService) GetFollowers(id string, offset int, limit int) ([]models.PersonDto, error) {
	if len(id) == 0 {
		return nil, errors.New("request is invalid")
	}
	persons, err := s.socRepo.GetFollowers(id, offset, limit)
	if err != nil {
		return nil, err
	}
	var result []models.PersonDto
	for _, person := range persons {
		result = append(result, models.PersonDto{
			Id:    person.Id,
			Name:  person.Name,
			Image: person.Image,
		})
	}
	return result, nil

}
