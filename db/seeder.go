package main

import (
	"fmt"
	"os"

	"github.com/arifin2018/facebook/db/seeder"
)

func main() {
	// Mendapatkan argumen baris perintah
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run seeder.go <function>")
		return
	}
	// Menjalankan fungsi berdasarkan argumen baris perintah
	switch args[1] {
	case "SeedRole":
		seeder.SeedRole()
	case "SeedPermission":
		seeder.SeedPermission()
	case "RolePermission":
		seeder.SeedRolePermission()
	default:
		fmt.Println("Unknown function")
	}
}
