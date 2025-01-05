package ArticleTagDao

import (
	"gorm.io/gorm"
)

type ArticleTag struct {
	Id        int    `db:"id"`
	ArticleId int    `db:"article_id"`
	TagId     int    `db:"tag_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}

func CreateArticleTag(db *gorm.DB, articleTag *ArticleTag) (bool, error) {
	tx := db.Create(&articleTag)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
func BatchCreateArticleTag(db *gorm.DB, articleTags []*ArticleTag) error {
	tx := db.Create(&articleTags)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func GetArticleTagByArticleId(db *gorm.DB, articleId int) ([]*ArticleTag, error) {
	var articleTag []*ArticleTag
	tx := db.Where(&ArticleTag{ArticleId: articleId}).Find(&articleTag)
	if tx.Error != nil {
		return articleTag, tx.Error
	}
	return articleTag, nil
}
