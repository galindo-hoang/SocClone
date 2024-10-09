package services

import "github.com/PostService/pkg/services/models"

type IPostService interface {
	CreatePost(post models.PostObject) (*models.PostRes, error)
	GetPost(id int) (*models.PostRes, error)
	UpdatePost(post models.MPostObject) (*models.PostRes, error)
	DeletePost(id int) error
	GetListPostFromWall(object models.ListPostUserObject) ([]*models.PostRes, error)
	LikePost(object models.LikePostObject) (int, error)
	UnLikePost(object models.LikePostObject) error
	// suggest using rgpc for real time
	CommentPost(object models.CommentPostObject) error
	DeleteComment(object models.CommentPostObject) error
}
