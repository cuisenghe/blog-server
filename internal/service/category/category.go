package category

import (
	"blog-server/internal/common/response"
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
func (s *service) GetCategoryList(ctx *gin.Context, tagName string) (*response.PageListResponse, error) {
	// dao
	var resp response.PageListResponse
	categorys, err := categoryDao.GetCategoryListByCondition(GetDB(ctx), tagName)
	if err != nil {
		return &resp, err
	}
	count, err := categoryDao.GetCategoryCountByCondition(GetDB(ctx), tagName)
	if err != nil {
		return &resp, err
	}
	resp.List = convertData(categorys)
	resp.Total = count
	return &resp, nil
}
func (s *service) AddCategory(ctx *gin.Context, tagName string) (*SimpleCategory, error) {
	var category categoryDao.Category
	category.CategoryName = tagName
	err := categoryDao.CreateCategory(GetDB(ctx), &category)
	if err != nil {
		return nil, err
	}
	return &SimpleCategory{
		ID:           category.ID,
		CategoryName: category.CategoryName,
	}, nil
}
