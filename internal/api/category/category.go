package category

import (
	"blog-server/internal/api"
	"blog-server/internal/common/constants"
	"blog-server/internal/common/response"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (h *Handler) GetCategoryDictionary(ctx *gin.Context) {
	dict, err := h.service.GetCategoryDict(ctx)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, dict)
}

type GetCategoryListReq struct {
	Current      int    `json:"current"`
	Size         int    `json:"size"`
	CategoryName string `json:"category_name"`
}
type AddCategoryReq struct {
	CategoryName string `json:"category_name"`
}

func (h *Handler) GetCategoryList(ctx *gin.Context) {
	var req GetCategoryListReq
	api.Binding(ctx, &req)
	// service
	resp, err := h.service.GetCategoryList(ctx, req.CategoryName)
	if err != nil {
		slog.Error(err.Error())
		response.Fail(ctx, constants.FAIL, "获取目录失败")
		return
	}
	response.SuccessWithPage(ctx, resp)
}
func (h *Handler) AddCategory(ctx *gin.Context) {
	var req AddCategoryReq
	api.Binding(ctx, &req)
	//
	resp, err := h.service.AddCategory(ctx, req.CategoryName)
	if err != nil {
		slog.Error(err.Error())
		response.Fail(ctx, constants.FAIL, "添加分类失败")
		return
	}
	response.Success(ctx, resp)
}
