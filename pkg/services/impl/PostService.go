package impl

import (
	"fmt"
	"github.com/PostService/pkg/repositories"
	modelsRepo "github.com/PostService/pkg/repositories/models"
	"github.com/PostService/pkg/services/models"
	"github.com/PostService/pkg/utils"
	"github.com/google/uuid"
	"path/filepath"
	"time"
)

type PostService struct {
	postRepo  repositories.IPostRepository
	cloudRepo repositories.ICloudStorageRepository
}

func NewPostService(post repositories.IPostRepository, cloud repositories.ICloudStorageRepository) *PostService {
	return &PostService{
		postRepo:  post,
		cloudRepo: cloud,
	}
}

func (s *PostService) CreatePost(model models.PostObject) (*models.PostRes, error) {
	var uid = uuid.New()
	var path = fmt.Sprintf("%v_%v/", model.Author, uid)
	defer func() {
		utils.DeleteDirectory(path)
	}()
	var files []modelsRepo.CloudEntity
	for _, file := range model.Files {
		path := filepath.Join(path, file.Filename)
		if err := utils.SaveFileToDes(file, path); err != nil {
			utils.DeleteDirectory(path)
			return nil, err
		} else {
			files = append(files, modelsRepo.CloudEntity{
				Retry:       3,
				FilePath:    path,
				Size:        file.Size,
				ObjectName:  file.Filename,
				ContentType: file.Header.Get("Content-Type"),
			})
		}
	}

	images, err := s.cloudRepo.UploadFiles(files, fmt.Sprintf("%vpost%v", model.Author, uid))
	if err != nil {
		return nil, err
	}

	var post = &modelsRepo.Posts{
		Description: model.Description,
		Author:      model.Author,
		CreateAt:    time.Now(),
		LastUpdated: time.Now(),
		Images:      images,
	}

	entity, err := s.postRepo.CreatePost(post)
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, image := range entity.Images {
		if len(image.Cdn) != 0 {
			urls = append(urls, image.Cdn)
		} else {
			urls = append(urls, image.S3)
		}
	}

	return &models.PostRes{
		ID:          entity.ID,
		Images:      urls,
		Author:      entity.Author,
		Likes:       len(entity.Likes),
		CreateAt:    entity.CreateAt,
		Description: entity.Description,
		LastUpdated: entity.LastUpdated,
	}, nil
}

func (s *PostService) GetPost(id int) (*models.PostRes, error) {
	entity, err := s.postRepo.FetchPostFromId(id)
	if err != nil {
		return nil, err
	}
	var urls []string
	println(entity.Images)
	for _, image := range entity.Images {
		if len(image.Cdn) != 0 {
			urls = append(urls, image.Cdn)
		} else {
			urls = append(urls, image.S3)
		}
	}

	return &models.PostRes{
		ID:          entity.ID,
		Images:      urls,
		Author:      entity.Author,
		Likes:       len(entity.Likes),
		CreateAt:    entity.CreateAt,
		Description: entity.Description,
		LastUpdated: entity.LastUpdated,
	}, nil
}

func (s *PostService) UpdatePost(post models.MPostObject) (*models.PostRes, error) {
	updatePost, err := s.postRepo.UpdatePost(post.PostId, modelsRepo.Posts{Description: *post.Description})
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, image := range updatePost.Images {
		if len(image.Cdn) != 0 {
			urls = append(urls, image.Cdn)
		} else {
			urls = append(urls, image.S3)
		}
	}

	return &models.PostRes{
		ID:          updatePost.ID,
		Images:      urls,
		Author:      updatePost.Author,
		Likes:       len(updatePost.Likes),
		CreateAt:    updatePost.CreateAt,
		Description: updatePost.Description,
		LastUpdated: updatePost.LastUpdated,
	}, nil
}

func (s *PostService) DeletePost(id int) error {
	return s.postRepo.DeletePost(id)
}

func (s *PostService) CommentPost(object models.CommentPostObject) error {
	var _, err = s.postRepo.CommentPost(&modelsRepo.Comments{
		From:     object.From,
		PostId:   object.PostId,
		Text:     object.Text,
		CreateAt: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *PostService) DeleteComment(object models.CommentPostObject) error {
	return s.postRepo.DeleteComment(&modelsRepo.Comments{ID: object.Id})
}

func (s *PostService) LikePost(object models.LikePostObject) (int, error) {
	return s.postRepo.CreateLike(&modelsRepo.Likes{PostId: object.PostId, AuthId: object.From})
}

func (s *PostService) UnLikePost(object models.LikePostObject) error {
	return s.postRepo.DeleteLike(&modelsRepo.Likes{PostId: object.PostId, AuthId: object.From})
}

func (s *PostService) GetListPostFromWall(object models.ListPostUserObject) ([]*models.PostRes, error) {
	var entities, err = s.postRepo.FetchListPostsFromUser(object.ID, object.Offset, object.Limit)
	if err != nil {
		return nil, err
	}
	var res = make([]*models.PostRes, len(entities))
	for _, entity := range entities {
		var images []string
		for _, image := range entity.Images {
			if len(image.Cdn) != 0 {
				images = append(images, image.Cdn)
			} else {
				images = append(images, image.S3)
			}
		}

		res = append(res, &models.PostRes{
			ID:          entity.ID,
			Likes:       len(entity.Likes),
			Author:      entity.Author,
			Images:      images,
			CreateAt:    entity.CreateAt,
			Description: entity.Description,
			LastUpdated: entity.LastUpdated,
		})
	}

	return res, nil
}
