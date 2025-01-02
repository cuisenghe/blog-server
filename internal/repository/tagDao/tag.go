package tagDao

import (
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
func getTagById(db *gorm.DB, id int) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}
func GetTagIdsByArticleId(db *gorm.DB, id int) ([]int, error) {
	var tagIds []int
	tx := db.Where("article_id = ?", id).Find(&tagIds)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tagIds, nil
}
func GetTagsById(db *gorm.DB, tagId []int) ([]*Tag, error) {
	var tags []*Tag
	tx := db.Where("id in (?)", tagId)
	tx.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}
