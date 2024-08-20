package repositories

import (
	"github.com/PostService/pkg/repositories/models"
)

type IPostRepository interface {
	FetchPostFromId(id int) (*models.Posts, error)
	CreatePost(posts *models.Posts) (*models.Posts, error)
	UpdatePost(id int, posts models.Posts) (*models.Posts, error)
	FetchListPostsFromUser(author int, offset int, limit int) ([]*models.Posts, error)
	DeletePost(id int) error

	CreateLike(posts *models.Like) (*models.Posts, error)

	// switch using realtime
	CommentPost(posts *models.Comment) (*models.Posts, error)
}
