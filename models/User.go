package models

type User struct {
	Id       string `gorm:"primaryKey;autoIncrement"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	Image    string `validate:"required" json:"image"`
}