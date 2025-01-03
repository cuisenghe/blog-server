package blogConfig

import (
	"blog-server/constants"
	"blog-server/internal/common/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetConfig(ctx *gin.Context) {
	config, err := h.service.GetConfig(ctx)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, config)
}
