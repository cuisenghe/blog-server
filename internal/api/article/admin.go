package article

import (
	"blog-server/internal/api"
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"blog-server/internal/repository/categoryDao"
	"blog-server/internal/service/article"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddArticleReq struct {
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
type TitleExistReq struct {
	ID           int    `json:"id"`
	ArticleTitle string `json:"article_title"`
}

func (h *Handler) AddArticle(ctx *gin.Context) {
	var req AddArticleReq
	api.Binding(ctx, &req)
	// 新增
	addArticle, err := h.service.AddArticle(ctx, convertData(&req))
	if err != nil {
		response.FailWithBizError(ctx, err)
	}
	response.Success(ctx, addArticle)
}
func convertData(req *AddArticleReq) *article.AddArticleData {
	// 封装tag
	tags := make([]*article.Tag, 0, len(req.TagList))
	for _, tag := range req.TagList {
		tags = append(tags, &article.Tag{
			ID:      tag.ID,
			TagName: tag.TagName,
		})
	}
	// 封装cover
	covers := make([]*article.Cover, 0, len(req.CoverList))
	for _, cover := range req.CoverList {
		covers = append(covers, &article.Cover{
			Name:       cover.Name,
			Percentage: cover.Percentage,
			Status:     cover.Status,
			Size:       cover.Size,
			Raw:        cover.Raw,
			Uid:        cover.Uid,
			Url:        cover.Url,
		})
	}
	return &article.AddArticleData{
		ArticleTitle: req.ArticleTitle,
		Category: &categoryDao.Category{
			ID:           req.Category.ID,
			CategoryName: req.Category.CategoryName,
		},
		TagList:            tags,
		AuthorID:           req.AuthorID,
		ArticleContent:     req.ArticleContent,
		ArticleCover:       req.ArticleCover,
		IsTop:              req.IsTop,
		Status:             req.Status,
		Type:               req.Type,
		OriginUrl:          req.OriginUrl,
		CoverList:          covers,
		ArticleDescription: req.ArticleDescription,
	}
}

// UpdateArticle 更新
func (h *Handler) UpdateArticle(ctx *gin.Context) {
	var req AddArticleReq
	api.Binding(ctx, &req)
	updateArticle, err := h.service.UpdateArticle(ctx, convertData(&req))
	if err != nil {
		response.Fail(ctx, 1, "获取错误")
	}
	response.Success(ctx, updateArticle)
}

// DeleteArticle 根据状态删除文章
func (h *Handler) DeleteArticle(ctx *gin.Context) {
	// binding
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, "获取错误")
	}
	status, err := strconv.Atoi(ctx.Param("status"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, "获取错误")
	}

	// service
	deleteArticle, err := h.service.DeleteArticle(ctx, id, status)
	if err != nil {
		response.Fail(ctx, constants.FAIL, "获取错误")
	}
	response.Success(ctx, deleteArticle)
}

// RevertArticle 恢复文章
func (h *Handler) RevertArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, "获取错误")
	}
	revertArticle, err := h.service.RevertArticle(ctx, id)
	if err != nil {
		response.Fail(ctx, constants.FAIL, "获取错误")
	}
	response.Success(ctx, revertArticle)
}

// TitleExist 判断文章是否存在
func (h *Handler) TitleExist(ctx *gin.Context) {
	var req TitleExistReq
	api.Binding(ctx, &req)
	// service
	exist, err := h.service.TitleExist(ctx, req.ID)
	if err != nil {
		response.Fail(ctx, constants.FAIL, "判断文章错误")
	}
	response.Success(ctx, exist)
}

// IsPublic 切换文章私密性
func (h *Handler) IsPublic(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	status, err := strconv.Atoi(ctx.Param("status"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	articleStatus, err := h.service.UpdateArticleStatus(ctx, id, status)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, articleStatus)
}

// UpdateTop 更改置顶
func (h *Handler) UpdateTop(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	is_top, err := strconv.Atoi(ctx.Param("is_top"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	// Service
	top, err := h.service.UpdateArticleTop(ctx, id, is_top)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, top)
}

// AdminGetArticleList 后台获取文章的列表
func (h *Handler) AdminGetArticleList(ctx *gin.Context) {
	var req *ListArticleReq
	api.Binding(ctx, &req)
	// Service
	resp, err := h.service.AdminGetArticleList(ctx, &article.ArticleListData{
		Current: req.Current,
		Size:    req.Size,
	})
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, resp)
}
