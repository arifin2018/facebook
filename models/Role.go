package models

import "gorm.io/gorm"

type Role struct {
	Name string `gorm:"column:name" validate:"required,role-unique-data-column=name"`
	gorm.Model
}
