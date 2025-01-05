package article

import (
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service interface {
	GetArticleList(ctx *gin.Context, req *ArticleListData) (*response.PageListResponse, error)

	AddArticle(ctx *gin.Context, data *AddArticleData) (bool, error)
	UpdateArticle(ctx *gin.Context, data *AddArticleData) (bool, error)
	DeleteArticle(ctx *gin.Context, id, status int) (bool, error)
	RevertArticle(ctx *gin.Context, id int) (bool, error)
	TitleExist(ctx *gin.Context, id int) (bool, error)
	UpdateArticleStatus(ctx *gin.Context, id, status int) (bool, error)

	UpdateArticleTop(ctx *gin.Context, id, is_top int) (bool, error)
	AdminGetArticleList(ctx *gin.Context, data *AdminArticleListData) (*response.PageListResponse, error)

	// 前端
	BlogTimelineGetArticleList(ctx *gin.Context, data *ArticleListData) (*response.PageListResponse, error)
	GetArticleListByTagId(ctx *gin.Context, data *ArticleListData) (*response.PageListResponse, error)
	GetArticleListByCategoryId(ctx *gin.Context, data *ArticleListData) (*response.PageListResponse, error)
	GetRecommendArticleById(ctx *gin.Context, articleId int) (*response.PageListResponse, error)
	GetArticleListByContent(ctx *gin.Context, content string) (*response.PageListResponse, error)
	GetHotArticle(ctx *gin.Context) (*response.PageListResponse, error)
	GetArticleById(ctx *gin.Context, articleId int) (*DetailArticle, error)
}
type service struct {
}

func NewService() Service {
	return &service{}
}
func GetDB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet(constants.DB).(*gorm.DB)
}
