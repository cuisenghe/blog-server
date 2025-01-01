package tagDao

import (
	"blog-server/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	ID        int
	TagName   string
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

func (Tag) TableName() string {
	return "blog_tag"
}
func getTagById(ctx *gin.Context, id int) (*Tag, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var tag Tag
	err := db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}
func GetTagIdsByArticleId(ctx *gin.Context, id int) ([]int, error) {
	var tagIds []int
	db := ctx.MustGet(constants.DB).(*gorm.DB)
	tx := db.Where("article_id = ?", id).Find(&tagIds)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tagIds, nil
}
func GetTagsById(ctx *gin.Context, tagId []int) ([]*Tag, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var tags []*Tag
	tx := db.Where("id in (?)", tagId)
	tx.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}
