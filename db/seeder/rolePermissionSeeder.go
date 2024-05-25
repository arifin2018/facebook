package seeder

import (
	"fmt"
	"log"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"gorm.io/gorm"
)

// go run db/seeder.go RolePermission
func roleAdmin(tx *gorm.DB, modelRolePermission *[]models.RolePermission) {
	roleAdmin := models.Role{
		Name: "admin",
	}
	tx.First(&roleAdmin, "name = ?", roleAdmin.Name)

	var permissionsAdmin = []string{
		"create_admin",
		"update_admin",
		"read_admin",
		"delete_admin",
	}
	modelPermissionAdmin := []models.Permission{}
	tx.Where("Name IN ?", permissionsAdmin).Find(&modelPermissionAdmin)
	for _, v := range modelPermissionAdmin {
		var valueModelPermissionAdmin = models.RolePermission{
			RoleId:       int(roleAdmin.ID),
			PermissionId: int(v.ID),
		}
		*modelRolePermission = append(*modelRolePermission, valueModelPermissionAdmin)
	}
}

func roleStaff(tx *gorm.DB, modelRolePermission *[]models.RolePermission) {

	roleStaff := models.Role{
		Name: "staff",
	}
	tx.First(&roleStaff, "name = ?", roleStaff.Name)

	var permissionsStaff = []string{
		"create_staff",
		"update_staff",
		"read_staff",
		"delete_staff",
	}
	modelPermissionStaff := []models.Permission{}
	tx.Where("Name IN ?", permissionsStaff).Find(&modelPermissionStaff)
	for _, v := range modelPermissionStaff {
		var valueModelPermissionStaff = models.RolePermission{
			RoleId:       int(roleStaff.ID),
			PermissionId: int(v.ID),
		}
		*modelRolePermission = append(*modelRolePermission, valueModelPermissionStaff)
	}
}
func roleUser(tx *gorm.DB, modelRolePermission *[]models.RolePermission) {
	roleUser := models.Role{
		Name: "user",
	}
	tx.First(&roleUser, "name = ?", roleUser.Name)

	var permissionsUser = []string{
		"create_user",
		"update_user",
		"read_user",
		"delete_user",
	}
	modelPermissionUser := []models.Permission{}
	tx.Where("Name IN ?", permissionsUser).Find(&modelPermissionUser)
	for _, v := range modelPermissionUser {
		var valueModelPermissionUser = models.RolePermission{
			RoleId:       int(roleUser.ID),
			PermissionId: int(v.ID),
		}
		*modelRolePermission = append(*modelRolePermission, valueModelPermissionUser)
	}

}

func SeedRolePermission() {
	tx := config.DB.Begin()
	modelRolePermission := []models.RolePermission{}
	roleAdmin(tx, &modelRolePermission)
	roleStaff(tx, &modelRolePermission)
	roleUser(tx, &modelRolePermission)

	if err := SaveRolePermissionsToDatabase(&modelRolePermission, tx); err != nil {
		log.Fatal("Failed to save user to database:", err)
	}
	tx.Commit()
}

// SaveUserToDatabase function untuk menyimpan data pengguna ke dalam database (contoh)
func SaveRolePermissionsToDatabase(rolePermission *[]models.RolePermission, tx *gorm.DB) error {
	// Simulasikan penyimpanan data ke dalam database
	// Di sini Anda bisa menggunakan ORM atau database driver yang digunakan dalam proyek Anda
	// Contoh sederhana: hanya mencetak informasi pengguna
	if err := config.DB.Create(rolePermission).Error; err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Saved user to database: %v\n", rolePermission)
	return nil
}
