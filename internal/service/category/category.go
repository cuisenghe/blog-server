package category

import (
	"blog-server/internal/repository/categoryDao"
	"github.com/gin-gonic/gin"
)

type SimpleCategory struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

func (s *service) GetCategoryDict(ctx *gin.Context) ([]*SimpleCategory, error) {
	category, err := categoryDao.GetCategory(GetDB(ctx))
	if err != nil {
		return nil, err
	}
	return convertData(category), nil
}
func convertData(data []*categoryDao.Category) []*SimpleCategory {
	var result []*SimpleCategory
	for _, category := range data {
		result = append(result, &SimpleCategory{
			ID:           category.ID,
			CategoryName: category.CategoryName,
		})
	}
	return result
}
