package models

import "gorm.io/gorm"

type Permission struct {
	ID   uint   `json:"-"`
	Name string `gorm:"column:name" validate:"required,permission-unique-data-column=name"`
	gorm.Model
}
