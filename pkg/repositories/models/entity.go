package models

import "time"

type Posts struct {
	ID          int        `gorm:"primaryKey; index; AUTO_INCREMENT"`
	Author      int        `gorm:"index"`
	CreateAt    time.Time  `gorm:"column:create_at"`
	LastUpdated time.Time  `gorm:"column:last_updated"`
	Images      []*Image   `gorm:"constraint:OnDelete:CASCADE;foreignKey:PostId;references:ID"`
	Comments    []*Comment `gorm:"constraint:OnDelete:CASCADE;foreignKey:PostId;references:ID"`
	Likes       []*Like    `gorm:"constraint:OnDelete:CASCADE;foreignKey:PostId;references:ID"`
	Description string
}

type Comment struct {
	ID       int       `gorm:"primaryKey; AUTO_INCREMENT"`
	From     string    `gorm:"column:id"`
	Text     string    `gorm:"text; not null"`
	CreateAt time.Time `gorm:"column:create_at"`
	PostId   int
}

type Image struct {
	ID     int    `gorm:"primaryKey; AUTO_INCREMENT"`
	Cdn    string `gorm:"column:cdn"`
	S3     string `gorm:"column:s3"`
	PostId int
}

type Like struct {
	From   string `gorm:"primaryKey; column:id"`
	PostId int    `gorm:"primaryKey; column:post_id"`
}
