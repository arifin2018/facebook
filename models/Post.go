package models

import (
	"math"
	"os"
	"strconv"

	"github.com/arifin2018/facebook/config"
	response "github.com/arifin2018/facebook/helpers/handlers/Response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Post struct {
	UserId  uint   `gorm:"column:User_id" validate:"required" json:"user_id"`
	Content string `gorm:"column:Content" validate:"required,check-userid-same-content-post" json:"content"`
	gorm.Model
}

func (p Post) DataPagination(f *fiber.Ctx, data *response.PaginationData) {
	config.DB.Model(p).Count(data.Count)
	config.DB.Scopes(response.Paginate(f, data.TotalPages)).Find(data.Model.(*[]Post))
	*data.CountTotalPages = math.Ceil(float64(*data.Count) / float64(*data.TotalPages))

	var nextLink string
	pageParams := f.Query("page")
	pageSize := f.Query("page_size")
	pageParamsConvertInt, _ := strconv.Atoi(pageParams)
	CountTotalPages := int(*data.CountTotalPages)
	if pageParamsConvertInt == CountTotalPages || pageParamsConvertInt > CountTotalPages {
		nextLink = strconv.Itoa(1)
	} else {
		nextLink = strconv.Itoa(pageParamsConvertInt + 1)
	}
	*data.Link = os.Getenv("BASE_URL") + "/api/post?page=" + nextLink + "&page_size=" + pageSize
}
