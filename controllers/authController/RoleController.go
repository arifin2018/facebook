package authcontroller

import (
	"log"
	"strconv"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/models/requestValidator"
	"github.com/arifin2018/facebook/services"
	"github.com/gofiber/fiber/v2"
)

func RoleIndex(f *fiber.Ctx) error {
	role := new([]models.Role)
	if err := services.IndexRole(f, role); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": role,
	})
}
func RoleCreate(f *fiber.Ctx) error {
	role := new(models.Role)
	if err := f.BodyParser(role); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := requestValidator.HandlersValidateStructRole(f, role); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := services.CreateRole(f, role); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": role,
	})
}
func RoleUpdate(f *fiber.Ctx) error {
	role := new(models.Role)
	id := f.Params("id")
	intId, _ := strconv.Atoi(id)
	if err := f.BodyParser(role); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := requestValidator.HandlersValidateStructRole(f, role); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := services.UpdateRole(f, intId, role); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": role,
	})
}
func RoleDelete(f *fiber.Ctx) error {
	id := f.Params("id")
	intId, _ := strconv.Atoi(id)
	if err := services.DeleteRole(f, intId); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": models.Permission{},
	})
}
