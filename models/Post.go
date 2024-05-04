package models

import "gorm.io/gorm"

type Post struct {
	Id      string `gorm:"primaryKey;autoIncrement"`
	UserId  uint   `validate:"required" json:"user_id"`
	Content string `validate:"required" json:"content"`
	gorm.Model
}
