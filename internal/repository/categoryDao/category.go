package categoryDao

import (
	"blog-server/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct {
	ID           int    `db:"id" json:"id"`
	CategoryName string `db:"category_name" json:"category_name"`
	CreatedAt    string `db:"createdAt" json:"createdAt"`
	UpdatedAt    string `db:"updatedAt" json:"updatedAt"`
}

func (Category) TableName() string {
	return "blog_category"
}

func GetCategory(ctx *gin.Context) ([]*Category, error) {
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	var categories []*Category
	tx := db.Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categories, nil
}
