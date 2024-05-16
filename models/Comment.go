package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	UserId  uint   `gorm:"column:User_id" json:"user_id"`
	PostId  uint   `gorm:"column:Post_id" validate:"required" json:"post_id"`
	Content string `gorm:"column:Content" validate:"required" json:"content"`
	gorm.Model
}
