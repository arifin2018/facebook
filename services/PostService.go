package services

import (
	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var postClass = repositories.PostClass{
	Post: models.Post{},
}

func GetPost(f *fiber.Ctx, posts *[]models.GetPost) {
	postClass.FindPost(f, posts)
}

func FindPostByID(f *fiber.Ctx, id int) models.Post {
	postClass.Post.ID = uint(id)
	postClass.FindByIDPost(f)
	return postClass.Post
}

func CreatePost(f *fiber.Ctx, post *models.Post) *models.Post {
	postClass.Post = models.Post{
		UserId:  uint(auth.MeData.User.Id),
		Content: post.Content,
	}
	postClass.CreatePost(f)
	return &postClass.Post
}

func UpdatePost(f *fiber.Ctx, post *models.Post) models.Post {
	postClass.Post.ID = post.ID
	postClass.FindByIDPost(f)
	postClass.Post.Content = post.Content
	postClass.UpdatePost(f)
	return postClass.Post
}

func DeletePost(f *fiber.Ctx, post *models.Post) {
	postClass.Post.ID = post.ID
	postClass.DeletePost(f)
}
