package article

import (
	"blog-server/internal/repository/articleDao"
	"github.com/gin-gonic/gin"
	"time"
)

type AddArticleData struct {
	ArticleTitle       string    `json:"article_title"`
	Category           *Category `json:"category"`
	TagList            []*Tag    `json:"tag_list"`
	AuthorID           int       `json:"author_id"`
	ArticleContent     string    `json:"article_content"`
	ArticleCover       string    `json:"article_cover"`
	IsTop              int       `json:"is_top"`
	Status             int       `json:"status"`
	Type               int       `json:"type"`
	OriginUrl          string    `json:"origin_url"`
	CoverList          []*Cover  `json:"cover_list"`
	ArticleDescription string    `json:"article_description"`
}
type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
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
	article, err := articleDao.CreateArticle(ctx, convertData(data))
	if err != nil {
		return false, err
	}
	// todo 创建tag
	// todo 创建category

	return article, nil
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
