package seeder

import (
	"fmt"
	"time"
)

// User struct untuk representasi data pengguna
type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
}

// SeedUser function untuk menambahkan data pengguna awal ke dalam database
func SeedUser() {
	// Simulasikan proses menyimpan data ke dalam database
	// user := User{
	// 	ID:        1,
	// 	Username:  "john_doe",
	// 	Email:     "john.doe@example.com",
	// 	CreatedAt: time.Now(),
	// }

	// // Contoh: Menyimpan data ke dalam database
	// if err := SaveUserToDatabase(user); err != nil {
	// 	log.Fatal("Failed to save user to database:", err)
	// }

	fmt.Println("Seeder: User data seeded successfully!")
}

// SaveUserToDatabase function untuk menyimpan data pengguna ke dalam database (contoh)
func SaveUserToDatabase(user User) error {
	// Simulasikan penyimpanan data ke dalam database
	// Di sini Anda bisa menggunakan ORM atau database driver yang digunakan dalam proyek Anda
	// Contoh sederhana: hanya mencetak informasi pengguna
	fmt.Printf("Saved user to database: %v\n", user)
	return nil
}
