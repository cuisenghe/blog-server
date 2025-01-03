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

func (h *Handler) GetTagDictionary(ctx *gin.Context) {

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
	}
	response.SuccessWithPage(ctx, resp)
}
