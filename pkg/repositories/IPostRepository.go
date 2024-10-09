package repositories

import (
	"github.com/PostService/pkg/repositories/models"
)

type IPostRepository interface {
	FetchPostFromId(id int) (*models.Posts, error)
	CreatePost(posts *models.Posts) (*models.Posts, error)
	UpdatePost(id int, posts models.Posts) (*models.Posts, error)
	//UpdateImage(posts models.Posts) (*models.Posts, error)
	FetchListPostsFromUser(author int, offset int, limit int) ([]*models.Posts, error)
	DeletePost(id int) error

	CreateLike(likes *models.Likes) (int, error)
	DeleteLike(posts *models.Likes) error

	// switch using realtime
	CommentPost(posts *models.Comments) (*models.Posts, error)
	DeleteComment(posts *models.Comments) error
}
