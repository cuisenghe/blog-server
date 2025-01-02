package categoryDao

import (
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
	tx := db.Create(category)
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
