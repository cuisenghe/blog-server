package category

import (
	"blog-server/constants"
	"blog-server/internal/common/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategoryDictionary(ctx *gin.Context) {
	dict, err := h.service.GetCategoryDict(ctx)
	if err != nil {
		response.Fail(ctx, constants.FAIL, err.Error())
	}
	response.Success(ctx, dict)
}
