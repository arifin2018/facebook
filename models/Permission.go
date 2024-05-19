package models

import "gorm.io/gorm"

type Permission struct {
	Name string `gorm:"column:name" validate:"required,permission-unique-data-column=name"`
	gorm.Model
}
