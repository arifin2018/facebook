package controllers

import (
	"log"
	"strconv"

	"github.com/arifin2018/facebook/helpers/auth"
	response "github.com/arifin2018/facebook/helpers/handlers/Response"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/models/requestValidator"
	"github.com/arifin2018/facebook/services"
	"github.com/gofiber/fiber/v2"
)

var postPaginate response.PaginateInterface = models.Post{}

func PostIndex(f *fiber.Ctx) error {
	posts := []models.GetPost{}
	services.GetPost(f, &posts)
	data := response.PaginationData{
		Count:           new(int64),
		TotalPages:      new(int64),
		CountTotalPages: new(float64),
		Model:           &posts,
		Link:            new(string),
	}
	// postPaginate.DataPagination(f, &data)
	// arifin
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data":              posts,
		"count_total_pages": data.CountTotalPages,
		"next_link":         data.Link,
	})
}

func PostFindById(f *fiber.Ctx) error {
	id, _ := strconv.Atoi(f.Params("id"))
	post := services.FindPostByID(f, id)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": post,
	})
}

func PostCreate(f *fiber.Ctx) error {
	post := new(models.Post)
	post.UserId = uint(auth.MeData.User.Id)
	f.BodyParser(post)
	if err := requestValidator.HandlersValidateStructPost(f, post); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusCreated).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	post = services.CreatePost(f, post)
	return f.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"data": post,
	})
}

func PostUpdate(f *fiber.Ctx) error {
	id, _ := strconv.Atoi(f.Params("id"))
	post := new(models.Post)
	post.UserId = uint(auth.MeData.User.Id)
	f.BodyParser(post)
	if err := requestValidator.HandlersValidateStructPost(f, post); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusCreated).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	post.ID = uint(id)
	postUpdate := services.UpdatePost(f, post)
	return f.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"data": postUpdate,
	})
}

func PostDelete(f *fiber.Ctx) error {
	id, _ := strconv.Atoi(f.Params("id"))
	post := new(models.Post)
	post.ID = uint(id)
	services.DeletePost(f, post)
	return f.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"data": post,
	})
}
