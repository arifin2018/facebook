package services

import (
	"strconv"

	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var postClass = repositories.PostClass{
	Post: models.Post{},
}

func GetPost(f *fiber.Ctx, posts *[]models.Post) {
	postClass.FindPost(f, posts)
}

func CreatePost(f *fiber.Ctx, post *models.Post) {
	userid, _ := strconv.Atoi(auth.MeData.User.Id)
	postClass.Post = models.Post{
		UserId:  uint(userid),
		Content: post.Content,
	}
	postClass.CreatePost(f)
}
