package category

import (
	"blog-server/internal/api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategoryDictionary(ctx *gin.Context) {
	dict, err := h.service.GetCategoryDict(ctx)
	if err != nil {
		api.ReturnBizError(ctx, err)
	}
	api.ReturnSuccess(ctx, dict)
}
