package impl

import (
	"errors"
	"github.com/PostService/pkg/database"
	"github.com/PostService/pkg/repositories"
	"github.com/PostService/pkg/repositories/models"
	"time"
)

type PostRepository struct{}

func NewPostRepository() repositories.IPostRepository {
	return &PostRepository{}
}

func (s *PostRepository) FetchPostFromId(id int) (*models.Posts, error) {
	var post *models.Posts = nil
	database.DB.Raw("SELECT * FROM posts WHERE ID = ?", id).Scan(&post)
	if post == nil {
		return nil, errors.New("post is invalid")
	}
	return post, nil
}

func (s *PostRepository) CreatePost(posts *models.Posts) (*models.Posts, error) {
	var result = database.DB.Create(posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (s *PostRepository) UpdatePost(id int, posts models.Posts) (*models.Posts, error) {
	result := database.DB.Raw("UPDATE posts SET description = ?, LastUpdated = ? where id = ?", posts.Description, time.Now(), id).Scan(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &posts, nil
}

func (s *PostRepository) FetchListPostsFromUser(author int, offset int, limit int) ([]*models.Posts, error) {
	var posts []*models.Posts = nil
	var result = database.DB.Raw("SELECT * FROM posts WHERE author = ? LIMIT ? OFFSET ?", author, limit, offset).Preload("Comments").Scan(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (s *PostRepository) DeletePost(id int) error {
	var user models.Posts
	result := database.DB.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *PostRepository) CreateLike(likes *models.Likes) (int, error) {
	var post *models.Posts = nil
	database.DB.Raw("SELECT * FROM posts WHERE id = ?", likes.PostId).Scan(&post)
	if post == nil {
		return 0, errors.New("post is invalid")
	}
	post.Likes = append(post.Likes, likes)
	result := database.DB.Save(&post)
	if result.Error != nil {
		return 0, result.Error
	}
	return len(post.Likes), nil
}

func (s *PostRepository) DeleteLike(likes *models.Likes) error {
	var res models.Posts
	result := database.DB.Delete(&res, likes.PostId, likes.AuthId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// switch using realtime
func (s *PostRepository) CommentPost(comment *models.Comments) (*models.Posts, error) {
	var post *models.Posts = nil
	database.DB.Raw("SELECT * FROM posts WHERE id = ?", comment.PostId).Scan(&post)
	if post == nil {
		return nil, errors.New("post is invalid")
	}
	post.Comments = append(post.Comments, comment)
	result := database.DB.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (s *PostRepository) DeleteComment(posts *models.Comments) error {
	var res models.Comments
	result := database.DB.Delete(&res, posts.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

//func (s *PostRepository) UpdateImage(posts models.Posts) (*models.Posts, error) {
//	return nil, nil
//}
