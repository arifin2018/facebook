package middlewares

import (
	"strings"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Di sini, Anda akan memeriksa peran pengguna dari sesi atau
		// data pengguna yang disimpan di konteks atau di mana pun yang sesuai
		// dengan kebutuhan aplikasi Anda.
		// Contoh sederhana: Anda bisa menyimpan peran pengguna di token JWT.
		// Anda kemudian akan memeriksa token JWT di sini dan mengizinkan
		// atau menolak akses berdasarkan peran pengguna.

		// Misalnya, Anda memiliki fungsi untuk memeriksa apakah pengguna memiliki peran yang sesuai
		// dengan rute yang diminta:
		if !userHasRole(c, role) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Access forbidden",
			})
		}
		// Jika pengguna memiliki peran yang sesuai, biarkan mereka melanjutkan ke rute.
		return c.Next()
	}
}

func userHasRole(c *fiber.Ctx, role string) bool {
	var result bool
	// Fungsi ini digunakan untuk memeriksa apakah pengguna memiliki peran yang sesuai.
	// Anda dapat mengimplementasikan logika sesuai dengan cara penyimpanan data pengguna Anda.
	// Misalnya, jika Anda menggunakan token JWT, Anda akan memeriksa token tersebut
	// dan mengekstrak peran pengguna dari token tersebut.
	// Di sini, kita akan mengasumsikan bahwa data pengguna tersedia di konteks Fiber,
	// Anda mungkin perlu menyesuaikannya sesuai dengan kebutuhan aplikasi Anda.
	user := c.Locals("user").(*jwt.Token)
	if user == nil {
		return false
	}
	// Lakukan pengecekan peran pengguna dan kembalikan hasilnya.
	// Ini akan tergantung pada bagaimana Anda mengatur data pengguna.
	// Misalnya, jika data pengguna adalah objek JSON, Anda akan mengekstrak peran dari objek tersebut.
	// Jika Anda menggunakan database, Anda mungkin telah menyimpan peran pengguna sebagai atribut di tabel pengguna.
	// Disesuaikan dengan cara Anda mengatur data pengguna.
	// Contoh sederhana:
	var dataUser = models.User{}
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	if err := config.DB.Model(&dataUser).Where("email = ?", email).Preload("UserRole").Preload("UserRole.RolePermisssion").First(&dataUser).Error; err != nil {
		return false
	}
	for _, v := range strings.Split(role, ",") {
		if dataUser.UserRole[0].Name == v {
			result = true
			break
		}
	}

	return result
}
