package models

import (
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	Image    string `validate:"required" json:"image"`
	Roleid   uint   `validate:"required" json:"roleid"`
	UserRole []Role `gorm:"many2many:user_roles;"`
}
type Me struct {
	User *User
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	userRole := UserRole{
		User_id: u.Id,
		Role_id: u.Roleid,
	}
	if err := tx.Create(&userRole).Error; err != nil {
		tx.Rollback()
		return err
	}
	return

}

func (u *User) TableName() string {
	return "users"
}
