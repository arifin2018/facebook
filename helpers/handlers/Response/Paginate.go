package response

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaginationData struct {
	Count           *int64
	TotalPages      *int64
	CountTotalPages *float64
	Model           interface{}
	Link			*string
}

type PaginateInterface interface {
	DataPagination(f *fiber.Ctx, data *PaginationData)
}

func Paginate(f *fiber.Ctx, total_pages *int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(f.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(f.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		*total_pages = int64(pageSize)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
