package impl

import (
	"github.com/SocService/pkg/repositories"
)

type SocService struct {
	socRepo *repositories.ISocRepo
}

func NewSocService(socRepo *repositories.ISocRepo) *SocService {
	return &SocService{socRepo: socRepo}
}

func (s *SocService) CreateUser()    {}
func (s *SocService) AddFollow()     {}
func (s *SocService) RemoveFollow()  {}
func (s *SocService) GetFollowings() {}
func (s *SocService) GetFollowers()  {}
