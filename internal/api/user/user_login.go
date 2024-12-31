package user

import (
	"blog-server/internal/api"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"authToken"`
}

func (h *Handler) Login(ctx *gin.Context) {
	var request LoginRequest
	api.Binding(ctx, &request)
	token, err := h.service.Login(ctx, request.Username, request.Password)
	if err != nil {
		api.ReturnBizError(ctx, err)
		return
	}
	api.ReturnSuccess(ctx, map[string]string{"token": token})
}
