package impl

import (
	"github.com/PostService/pkg/repositories"
	"github.com/PostService/pkg/services/models"
)

type PostService struct {
	repo repositories.IPostRepository
}

func NewPostService(repository repositories.IPostRepository) *PostService {
	return &PostService{}
}

func (s *PostService) CreatePost(post models.PostObject) (*models.PostRes, error) {
	return nil, nil
}

func (s *PostService) GetPost(id string) (*models.PostRes, error) {
	return nil, nil
}

func (s *PostService) UpdatePost(post models.MPostObject) (*models.PostRes, error) {
	return nil, nil
}

func (s *PostService) DeletePost(id string) (*models.PostRes, error) {
	return nil, nil
}

func (s *PostService) CommentPost(object models.CommentPostObject) error {
	return nil
}

func (s *PostService) LikePost(object models.LikePostObject) error {
	return nil
}

func (s *PostService) GetListPostFromWall(object models.ListPostUserObject) ([]*models.PostRes, error) {
	return nil, nil
}
