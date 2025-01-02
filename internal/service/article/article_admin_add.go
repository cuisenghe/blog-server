package article

import (
	"blog-server/internal/repository/ArticleTagDao"
	"blog-server/internal/repository/articleDao"
	"blog-server/internal/repository/categoryDao"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type AddArticleData struct {
	ArticleTitle       string                `json:"article_title"`
	Category           *categoryDao.Category `json:"category"`
	TagList            []*Tag                `json:"tag_list"`
	AuthorID           int                   `json:"author_id"`
	ArticleContent     string                `json:"article_content"`
	ArticleCover       string                `json:"article_cover"`
	IsTop              int                   `json:"is_top"`
	Status             int                   `json:"status"`
	Type               int                   `json:"type"`
	OriginUrl          string                `json:"origin_url"`
	CoverList          []*Cover              `json:"cover_list"`
	ArticleDescription string                `json:"article_description"`
}
type Tag struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
}
type Cover struct {
	Name       string      `json:"name"`
	Percentage int         `json:"percentage"`
	Status     string      `json:"status"`
	Size       int         `json:"size"`
	Raw        interface{} `json:"raw"`
	Uid        int         `json:"uid"`
	Url        string      `json:"url"`
}

func (s *service) AddArticle(ctx *gin.Context, data *AddArticleData) (bool, error) {
	// 创建文章
	// 开启事务
	tx := GetDB(ctx).Begin()
	article, err := articleDao.CreateArticle(tx, convertData(data))
	if err != nil {
		return false, err
	}
	// 判断目录是否存在，不存在就创建
	_, err = categoryDao.GetCategoryById(GetDB(ctx), data.Category.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 创建目录
		err := categoryDao.CreateCategory(GetDB(ctx), data.Category)
		if err != nil {
			return false, err
		}
	}
	// todo 创建tag
	var articleTags []*ArticleTagDao.ArticleTag
	for _, tag := range data.TagList {
		articleTags = append(articleTags, &ArticleTagDao.ArticleTag{
			TagId:     tag.ID,
			ArticleId: article.ID,
		})
	}
	err = ArticleTagDao.BatchCreateArticleTag(GetDB(ctx), articleTags)
	if err != nil {
		return false, err
	}
	// todo 创建category

	return true, nil
}
func convertData(data *AddArticleData) *articleDao.Article {
	return &articleDao.Article{
		ArticleTitle:       data.ArticleTitle,
		AuthorId:           data.AuthorID,
		CategoryId:         data.Category.ID,
		ArticleContent:     data.ArticleContent,
		ArticleCover:       data.ArticleCover,
		IsTop:              data.IsTop,
		Status:             data.Status,
		OriginUrl:          data.OriginUrl,
		CreateAt:           time.Now(),
		UpdateAt:           time.Now(),
		ArticleDescription: data.ArticleDescription,
	}
}

// UpdateArticle 更新
func (s *service) UpdateArticle(ctx *gin.Context, data *AddArticleData) (bool, error) {
	return true, nil
}
