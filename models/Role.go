package models

import "gorm.io/gorm"

type Role struct {
	ID              uint         `json:"-"`
	Name            string       `gorm:"column:name" validate:"required,role-unique-data-column=name"`
	RolePermisssion []Permission `gorm:"many2many:role_permisssions;foreignKey:ID;joinForeignKey:role_id;References:ID;joinReferences:permissions_id"`
	gorm.Model
}
