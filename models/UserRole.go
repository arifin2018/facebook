package models

type UserRole struct {
	User_id int `gorm:"column:user_id" validate:"required"`
	Role_id int `gorm:"column:role_id" validate:"required"`
}
