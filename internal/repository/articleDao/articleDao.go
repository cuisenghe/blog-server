package articleDao

import (
	constants2 "blog-server/internal/common/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID                 int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ArticleTitle       string    `gorm:"column:article_title" json:"article_title"`
	AuthorId           int       `gorm:"column:author_id" json:"author_id"`
	CategoryId         int       `gorm:"column:category_id" json:"category_id"`
	ArticleContent     string    `gorm:"column:article_content" json:"article_content"`
	ArticleCover       string    `gorm:"column:article_cover" json:"article_cover"`
	IsTop              int       `gorm:"column:is_top" json:"is_top"`
	Status             int       `gorm:"column:status" json:"status"`
	OriginUrl          string    `gorm:"column:origin_url" json:"origin_url"`
	CreateAt           time.Time `gorm:"column:createdAt" json:"create_at"`
	UpdateAt           time.Time `gorm:"column:updatedAt" json:"update_at"`
	ArticleDescription string    `gorm:"column:article_description" json:"article_description"`
	ThumbsUpTimes      uint      `gorm:"column:thumbs_up_times" json:"thumbs_up_times"`
	ReadingDuration    uint      `gorm:"column:reading_duration" json:"reading_duration"`
	Order              int       `gorm:"column:order" json:"order"`
}

func (a Article) TableName() string {
	return "blog_article"
}
func GetArticleList(db *gorm.DB, current, size int) ([]*Article, error) {
	var articles []*Article
	// 查询
	offset := (current - 1) * size
	tx := db.Limit(size).Offset(offset).Order("createdAt desc").Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}
func GetSumCount(db *gorm.DB) (int, error) {
	var count int64
	t := db.Model(&Article{}).Count(&count)
	if t.Error != nil {
		return 0, t.Error
	}
	return int(count), nil
}
func CreateArticle(db *gorm.DB, article *Article) (*Article, error) {
	tx := db.Create(&article)
	if tx.Error != nil {
		return article, tx.Error
	}
	return article, nil
}
func DeleteArticle(db *gorm.DB, id, status int) (bool, error) {
	tx := db.Where("id = ?", id).Update("status", status)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
func RevertArticle(db *gorm.DB, id int) (bool, error) {
	tx := db.Where("article_id = ?", id).Update("status", constants2.ARTICLE_PUBLIC_STATUS)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
func GetArticleById(db *gorm.DB, id int) (*Article, error) {
	var article Article
	tx := db.Where("id = ?", id).Find(&article)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &article, nil
}
func UpdateArticleStatus(ctx *gin.Context, id int, status int) (bool, error) {
	db := ctx.MustGet(constants2.DB).(*gorm.DB)
	tx := db.Model(&Article{}).Where("id = ?", id).Update("status", status)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
func UpdateArticleTop(db *gorm.DB, id int, is_top int) (bool, error) {
	tx := db.Model(&Article{}).Where("id = ?", id).Update("is_top", is_top)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
func GetArticleListByCondition(db *gorm.DB, current, size int, condition map[string]interface{}) ([]*Article, error) {
	//db
	var articles []*Article
	// 查询
	offset := (current - 1) * size
	// 判断condition是否为空
	if len(condition) > 0 {
		db = db.Where(condition)
	}
	tx := db.Limit(size).Offset(offset).Order("createAt desc").Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}
func GetArticleListByContent(db *gorm.DB, content string) ([]*Article, error) {
	var articles []*Article
	tx := db.Where("article_content Like ?", content).
		Where("status = ?", constants2.ARTICLE_PUBLIC_STATUS).
		Order("view_times desc").Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}
func GetArticleCountByContent(db *gorm.DB, content string) (int, error) {
	var count int64
	tx := db.Model(&Article{}).Where("article_content Like ?", content).
		Where("status = ?", constants2.ARTICLE_PUBLIC_STATUS).
		Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}
