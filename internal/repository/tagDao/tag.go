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
func BatchGetTags(db *gorm.DB, tagIds []int) ([]*Tag, error) {
	var tags []*Tag
	tx := db.Where("id in (?)", tagIds)
	tx.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}
func GetTagListByCondition(db *gorm.DB, tagName string) ([]*Tag, error) {
	var tags []*Tag
	if len(tagName) != 0 {
		db = db.Where("tag_name LIKE ?", "%"+tagName+"%")
	}
	tx := db.Model(&Tag{}).Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}
func GetCountByCondition(db *gorm.DB, tagName string) (int, error) {
	var count int64
	if len(tagName) != 0 {
		db = db.Where("tag_name LIKE ?", "%"+tagName+"%")
	}
	tx := db.Model(&Tag{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}
func AddTag(db *gorm.DB, tagName string) (*Tag, error) {
	var tag Tag
	tag.TagName = tagName
	tx := db.Create(&tag)
	if tx.Error != nil {
		return &tag, tx.Error
	}
	return &tag, nil
}
