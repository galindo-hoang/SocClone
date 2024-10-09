package repositories

import (
	model "github.com/SocService/pkg/repositories/model"
)

type ISocRepository interface {
	MakeRelation(from string, to string) error
	MakeDetach(from string, to string) error
	CreateNode(person model.Person) error
	GetFollowings(id string, offset int, limit int) ([]*model.Person, error)
	GetFollowers(id string, limit int, offset int) ([]*model.Person, error)
}
