package user

import (
	"blog-server/internal/api"
	"blog-server/internal/service/user"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" `
	QQ       string `json:"qq" `
	Avatar   string `json:"avatar" `
	Role     int    `json:"role" `
}
type RegisterResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) Register(ctx *gin.Context) {
	// 参数绑定
	var req RegisterRequest
	api.Binding(ctx, &req)

	// 走service
	id, err := h.service.Register(ctx, &user.RegData{
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
		QQ:       req.QQ,
		Avatar:   req.Avatar,
		Role:     req.Role,
	})
	if err != nil {
		api.ReturnBizError(ctx, err)
		return
	}
	api.ReturnSuccess(ctx, &RegisterResponse{
		ID:       id,
		Username: req.Username,
	})
}
