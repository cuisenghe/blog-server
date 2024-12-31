package blogConfig

import (
	"blog-server/internal/api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetConfig(ctx *gin.Context) {
	config, err := h.service.GetConfig(ctx)
	if err != nil {
		api.ReturnBizError(ctx, err)
	}
	api.ReturnSuccess(ctx, config)
}
