package article

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetArticleList(ctx *gin.Context, req *ArticleListData) (*ArticleListResp, error)

	AddArticle(ctx *gin.Context, data *AddArticleData) (bool, error)
	UpdateArticle(ctx *gin.Context, data *AddArticleData) (bool, error)
	DeleteArticle(ctx *gin.Context, id, status int) (bool, error)
	RevertArticle(ctx *gin.Context, id int) (bool, error)
	TitleExist(ctx *gin.Context, id int) (bool, error)
	UpdateArticleStatus(ctx *gin.Context, id, status int) (bool, error)

	UpdateArticleTop(ctx *gin.Context, id, is_top int) (bool, error)
	AdminGetArticleList(ctx *gin.Context, data *ArticleListData) (*ArticleListResp, error)

	// 前端
	BlogTimelineGetArticleList(ctx *gin.Context, data *ArticleListData) (*SimpleArticleListResp, error)
	GetArticleListByTagId(ctx *gin.Context, data *ArticleListData) (*SimpleArticleListResp, error)
	GetArticleListByCategoryId(ctx *gin.Context, data *ArticleListData) (*SimpleArticleListResp, error)
	GetRecommendArticleById(ctx *gin.Context, articleId int) (*RecommendArticleListResp, error)
	GetArticleListByContent(ctx *gin.Context, content string) (*ContentArticleListResp, error)
	GetHotArticle(ctx *gin.Context) (*SimpleArticleListResp, error)
	GetArticleById(ctx *gin.Context, articleId int) (*DetailArticle, error)
}
type service struct {
}

func NewService() Service {
	return &service{}
}
