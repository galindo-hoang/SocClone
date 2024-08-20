package models

import (
	"mime/multipart"
	"time"
)

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type PostObject struct {
	Description string                  `form:"description"`
	Author      string                  `form:"author"`
	Files       []*multipart.FileHeader `form:"files"`
}

type MPostObject struct {
	PostId      string  `json:"postId"`
	Author      string  `json:"author"`
	Description *string `json:"description"`
}

type LikePostObject struct {
	From   string `json:"id"`
	PostId string `json:"post_id"`
}

type CommentPostObject struct {
	From   string `json:"id"`
	PostId string `json:"post_id"`
}

type PostRes struct {
	ID          string    `json:"id"`
	Likes       int       `json:"likes"`
	Author      string    `json:"author"`
	Images      []string  `json:"images"`
	CreateAt    time.Time `json:"create_at"`
	Description string    `json:"description"`
	LastUpdated time.Time `json:"last_updated"`
}

type ListPostUserObject struct {
	ID     string `json:"id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
