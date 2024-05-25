package seeder

import (
	"fmt"
	"log"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"gorm.io/gorm"
)

// go run db/seeder.go SeedRole
func SeedRole() {
	tx := config.DB.Begin()
	var role = []models.Role{
		{Name: "admin"},
		{Name: "staff"},
		{Name: "user"},
	}
	if err := SaveRolesToDatabase(&role, tx); err != nil {
		log.Fatal("Failed to save user to database:", err)
	}
	tx.Commit()
}

// SaveUserToDatabase function untuk menyimpan data pengguna ke dalam database (contoh)
func SaveRolesToDatabase(role *[]models.Role, tx *gorm.DB) error {
	// Simulasikan penyimpanan data ke dalam database
	// Di sini Anda bisa menggunakan ORM atau database driver yang digunakan dalam proyek Anda
	// Contoh sederhana: hanya mencetak informasi pengguna
	if err := config.DB.Create(role).Error; err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Saved user to database: %v\n", role)
	return nil
}
