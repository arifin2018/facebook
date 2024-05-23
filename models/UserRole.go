package models

type UserRole struct {
	User_id uint `gorm:"column:user_id" validate:"required"`
	Role_id uint `gorm:"column:role_id" validate:"required"`
}
