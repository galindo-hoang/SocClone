package models

import "time"

type Posts struct {
	ID          int         `gorm:"primaryKey; index; AUTO_INCREMENT"`
	Author      string      `gorm:"index"`
	CreateAt    time.Time   `gorm:"column:create_at"`
	LastUpdated time.Time   `gorm:"column:last_updated"`
	Images      []*Images   `gorm:"constraint:OnDelete:CASCADE;foreignKey:PostId;references:ID"`
	Comments    []*Comments `gorm:"constraint:OnDelete:CASCADE;foreignKey:PostId;references:ID"`
	Likes       []*Likes    `gorm:"constraint:OnDelete:CASCADE;foreignKey:PostId;references:ID"`
	Description string
}

type Comments struct {
	ID       int       `gorm:"primaryKey; AUTO_INCREMENT"`
	From     string    `gorm:"column:id"`
	Text     string    `gorm:"text; not null"`
	CreateAt time.Time `gorm:"column:create_at"`
	PostId   int
}

type Images struct {
	ID     int    `gorm:"primaryKey; AUTO_INCREMENT"`
	Cdn    string `gorm:"column:cdn"`
	S3     string `gorm:"column:s3"`
	PostId int
}

type Likes struct {
	AuthId string `gorm:"primaryKey; column:id"`
	PostId int    `gorm:"primaryKey; column:post_id"`
}
