package categoryDao

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID           int       `db:"id" json:"id"`
	CategoryName string    `db:"category_name" json:"category_name"`
	CreatedAt    time.Time `db:"createdAt" json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt    time.Time `db:"updatedAt" json:"updatedAt" gorm:"column:updatedAt"`
}

func (Category) TableName() string {
	return "blog_category"
}

func GetCategory(db *gorm.DB) ([]*Category, error) {
	var categories []*Category
	tx := db.Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categories, nil
}
func GetCategoryById(db *gorm.DB, id int) (*Category, error) {
	var category Category
	tx := db.Where("id = ?", id).Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &category, nil
}
func CreateCategory(db *gorm.DB, category *Category) error {
	tx := db.FirstOrCreate(category, &Category{
		CategoryName: category.CategoryName,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func UpsertCategory(db *gorm.DB, category *Category) error {
	tx := db.Save(category)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func GetCategoryListByCondition(db *gorm.DB, categoryName string) ([]*Category, error) {
	var categories []*Category
	if len(categoryName) != 0 {
		db.Where("category_name like ?", "%"+categoryName+"%")
	}
	tx := db.Model(&Category{}).Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categories, nil
}
func GetCategoryCountByCondition(db *gorm.DB, categoryName string) (int, error) {
	var count int64
	if len(categoryName) != 0 {
		db.Where("category_name like ?", "%"+categoryName+"%")

	}
	tx := db.Model(&Category{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}
