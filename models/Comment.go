package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	UserId  uint   `gorm:"column:User_id" json:"userId"`
	PostId  uint   `gorm:"column:Post_id" validate:"required" json:"postId"`
	Content string `gorm:"column:Content" validate:"required" json:"content"`
	gorm.Model
}
