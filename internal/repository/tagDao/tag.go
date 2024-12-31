package tagDao

import (
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

func getTagById(ctx *gin.Context, id int) (*Tag, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var tag Tag
	err := db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}
