package authcontroller

import (
	"log"
	"strconv"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/models/requestValidator"
	"github.com/arifin2018/facebook/services"
	"github.com/gofiber/fiber/v2"
)

func PermissionIndex(f *fiber.Ctx) error {
	permissions := new([]models.Permission)
	if err := services.IndexPermission(f, permissions); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": permissions,
	})
}

func PermissionCreate(f *fiber.Ctx) error {
	permission := new(models.Permission)
	if err := f.BodyParser(permission); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := requestValidator.HandlersValidateStructPermission(f, permission); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := services.CreatePermission(f, permission); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": permission,
	})
}

func PermissionUpdate(f *fiber.Ctx) error {
	permission := new(models.Permission)
	id := f.Params("id")
	intId, _ := strconv.Atoi(id)
	if err := f.BodyParser(permission); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := requestValidator.HandlersValidateStructPermission(f, permission); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := services.UpdatePermission(f, intId, permission); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": permission,
	})
}

func PermissionDelete(f *fiber.Ctx) error {
	id := f.Params("id")
	intId, _ := strconv.Atoi(id)
	if err := services.DeletePermission(f, intId); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": models.Permission{},
	})
}
