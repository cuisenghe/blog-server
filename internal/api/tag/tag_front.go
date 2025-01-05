package tag

import (
	"blog-server/internal/api"
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"blog-server/internal/service/tag"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// tag前端需要的

type Front struct{}

type GetTagListReq struct {
	Current int    `json:"current" form:"current"`
	Size    int    `json:"size" form:"size"`
	TagName string `json:"tag_name" form:"tag_name"`
}
type AddTagReq struct {
	Id        string `json:"id" form:"id"`
	IsTrusted bool   `json:"isTrusted" form:"isTrusted"`
	TagName   string `json:"tag_name" form:"tag_name"`
	VTS       int    `json:"_vts" form:"_vts"`
}

func (h *Handler) GetTagDictionary(ctx *gin.Context) {
	resp, err := h.service.GetTagDict(ctx)
	if err != nil {
		slog.Error(err.Error())
		response.Fail(ctx, constants.FAIL, "获取Tag字典失败")
		return
	}
	response.Success(ctx, resp)
}

func (h *Handler) GetTagList(ctx *gin.Context) {
	var req GetTagListReq
	api.Binding(ctx, &req)
	// service
	resp, err := h.service.GetTagList(ctx, &tag.ListData{
		Current: req.Current,
		Size:    req.Size,
		TagName: req.TagName,
	})
	if err != nil {
		slog.Error(err.Error())
		response.Fail(ctx, constants.FAIL, err.Error())
		return
	}
	response.SuccessWithPage(ctx, resp)
}
func (h *Handler) AddTag(ctx *gin.Context) {
	var req AddTagReq
	api.Binding(ctx, &req)
	id, err := h.service.AddTag(ctx, req.TagName)
	if err != nil {
		slog.Error(err.Error())
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, gin.H{
		"id":       id,
		"tag_name": req.TagName,
	})
}
