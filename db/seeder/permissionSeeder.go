package seeder

import (
	"fmt"
	"log"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"gorm.io/gorm"
)

// go run db/seeder.go SeedRole
func SeedPermission() {
	tx := config.DB.Begin()
	var permissions = []models.Permission{
		{Name: "create_admin"},
		{Name: "update_admin"},
		{Name: "read_admin"},
		{Name: "delete_admin"},
		{Name: "create_staff"},
		{Name: "update_staff"},
		{Name: "read_staff"},
		{Name: "delete_staff"},
		{Name: "create_user"},
		{Name: "update_user"},
		{Name: "read_user"},
		{Name: "delete_user"},
	}
	if err := SavePermissionsToDatabase(&permissions, tx); err != nil {
		log.Fatal("Failed to save user to database:", err)
	}
	tx.Commit()
}

// SaveUserToDatabase function untuk menyimpan data pengguna ke dalam database (contoh)
func SavePermissionsToDatabase(Permission *[]models.Permission, tx *gorm.DB) error {
	// Simulasikan penyimpanan data ke dalam database
	// Di sini Anda bisa menggunakan ORM atau database driver yang digunakan dalam proyek Anda
	// Contoh sederhana: hanya mencetak informasi pengguna
	if err := config.DB.Create(Permission).Error; err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Saved user to database: %v\n", Permission)
	return nil
}
