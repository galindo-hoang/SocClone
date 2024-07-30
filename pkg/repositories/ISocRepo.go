package repositories

import (
	model "github.com/SocService/pkg/repositories/model"
)

type ISocRepo interface {
	MakeRelation(from string, to string) error
	MakeDetach(from string, to string) error
	CreateNode(person model.Person) error
	GetListRelationsFrom(person model.Person, offset int, limit int) ([]*model.Person, error)
	GetFromListRelations(person model.Person, limit int, offset int) ([]*model.Person, error)
}
