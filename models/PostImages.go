package models

import (
	"gorm.io/gorm"
)

type PostImages struct {
	PostId  uint   `gorm:"column:Post_id" validate:"required" json:"post_id"`
	Url 	string `gorm:"column:Url" validate:"required" json:"url"`
	Caption string `gorm:"column:Caption" validate:"required" json:"caption"`
	gorm.Model
}

