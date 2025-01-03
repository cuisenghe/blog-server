package article

import (
	"blog-server/internal/api"
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"blog-server/internal/service/article"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"

	"strconv"
)

type ListArticleReq struct {
	Size    int `json:"size"`
	Current int `json:"current"`
	Id      int `json:"id"`
}

// GetArticleList 分页获取主页文章内容
func (h *Handler) GetArticleList(ctx *gin.Context) {
	// binding
	size, err := strconv.Atoi(ctx.Param("size"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
		return
	}
	current, err := strconv.Atoi(ctx.Param("current"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
		return
	}
	// service
	resp, err := h.service.GetArticleList(ctx, &article.ArticleListData{
		Size:    size,
		Current: current,
	})
	if err != nil {
		response.Fail(ctx, constants.FAIL, "获取失败")
		return
	}
	response.SuccessWithPage(ctx, resp)
}

// 获取文章时间轴
func (h *Handler) BlogTimelineGetArticleList(ctx *gin.Context) {
	//
	current, err := strconv.Atoi(ctx.Param("current"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	size, err := strconv.Atoi(ctx.Param("size"))
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	// service
	resp, err := h.service.BlogTimelineGetArticleList(ctx, &article.ArticleListData{
		Size:    size,
		Current: current,
	})
	if err != nil {
		slog.Error("获取时间线失败：", err.Error())
		response.Fail(ctx, constants.FAIL, err.Error())
		return
	}

	response.SuccessWithPage(ctx, resp)
}
func (h *Handler) GetArticleListByTagId(ctx *gin.Context) {
	// 获取简略信息
	var req ListArticleReq
	api.Binding(ctx, &req)
	// service
	resp, err := h.service.GetArticleListByTagId(ctx, &article.ArticleListData{
		Current: req.Current,
		Size:    req.Size,
		Id:      req.Id,
	})
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.SuccessWithPage(ctx, resp)
}
func (h *Handler) GetArticleListByCategoryId(ctx *gin.Context) {
	var req ListArticleReq
	api.Binding(ctx, &req)
	// service
	resp, err := h.service.GetArticleListByCategoryId(ctx, &article.ArticleListData{
		Current: req.Current,
		Size:    req.Size,
		Id:      req.Id,
	})
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.SuccessWithPage(ctx, resp)
}
func (h *Handler) GetRecommendArticleById(ctx *gin.Context) {
	article_id := ctx.Param("article_id")
	fmt.Println(article_id)
}
func (h *Handler) GetArticleListByContent(ctx *gin.Context) {
	content := ctx.Param("content")
	resp, err := h.service.GetArticleListByContent(ctx, content)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.SuccessWithPage(ctx, resp)
}
func (h *Handler) GetHotArticle(ctx *gin.Context) {

}

func (h *Handler) GetArticleById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	resp, err := h.service.GetArticleById(ctx, id)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
		return
	}
	response.Success(ctx, resp)
}
